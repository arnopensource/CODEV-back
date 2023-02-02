package ade

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	_ent "github.com/abc3354/CODEV-back/ent"
	"github.com/abc3354/CODEV-back/ent/room"
	"github.com/abc3354/CODEV-back/services/ent"

	ics "github.com/arran4/golang-ical"
)

const fetchPeriod = 7 * 24 * time.Hour

type Class struct {
	start time.Time
	end   time.Time
}

type Room struct {
	name    string
	classes []Class
}

func Update() {
	updateStart := time.Now()

	downloadURL := generateURL(updateStart)

	httpClient := http.Client{
		Timeout: 5 * time.Minute,
	}
	res, err := httpClient.Get(downloadURL)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = res.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	rooms := parseICAL(res.Body)
	bounds := [2]time.Time{updateStart, updateStart.Add(fetchPeriod)}

	for _, adeRoom := range rooms {
		calculateAvailability(adeRoom, bounds)
	}
}

func calculateAvailability(adeRoom Room, bounds [2]time.Time) {
	currentTime := bounds[0]
	classPointer := 0
	dbRoom := ent.Get().Room.Query().Where(room.NameEQ(adeRoom.name)).OnlyX(context.Background())

	for currentTime.Before(bounds[1]) {
		switch {
		// not open
		case currentTime.Hour() < 8 || currentTime.Hour() >= 20:
			currentTime = currentTime.Add(1 * time.Hour)
		case currentTime.Weekday() == time.Saturday || currentTime.Weekday() == time.Sunday:
			currentTime = currentTime.Add(24 * time.Hour)

		// no more classes, available until 20h, it will jump to next day at next iteration
		case classPointer >= len(adeRoom.classes):
			from := currentTime
			currentTime = goto20(currentTime)
			saveAvailability(from, currentTime, dbRoom)

		// jump after class
		case adeRoom.classes[classPointer].start.Before(currentTime) || adeRoom.classes[classPointer].start.Equal(currentTime):
			end := adeRoom.classes[classPointer].end
			classPointer++
			if classPointer >= len(adeRoom.classes) {
				continue
			}
			if currentTime.Before(end) {
				currentTime = adeRoom.classes[classPointer].end
			}

		// next class is next day, available until 20h
		case adeRoom.classes[classPointer].start.Day() != currentTime.Day():
			from := currentTime
			currentTime = goto20(currentTime)
			saveAvailability(from, currentTime, dbRoom)

		// available until next class
		default:
			from := currentTime
			currentTime = adeRoom.classes[classPointer].start
			saveAvailability(from, currentTime, dbRoom)
		}
	}
}

func generateURL(start time.Time) string {
	res := url.URL{
		Scheme: "http",
		Host:   "adelb.univ-lyon1.fr",
		Path:   "/jsp/custom/modules/plannings/anonymous_cal.jsp",
	}

	values := res.Query()
	values.Set("resources", "49505,15954,49508,49507")
	values.Set("projectId", "2")
	values.Set("calType", "ical")
	values.Set("firstDate", start.Format("2006-01-02"))
	values.Set("lastDate", start.Add(fetchPeriod).Format("2006-01-02"))
	res.RawQuery = values.Encode()

	return res.String()
}

func parseICAL(reader io.Reader) map[string]Room {
	calendar, err := ics.ParseCalendar(reader)
	if err != nil {
		panic(err)
	}

	result := make(map[string]Room, 0)
	for _, component := range calendar.Components {
		event := component.(*ics.VEvent)
		properties := event.ComponentBase.Properties
		propertyMap := make(map[string]string)
		for _, property := range properties {
			propertyMap[property.IANAToken] = property.Value
		}

		location, ok := propertyMap["LOCATION"]
		if !ok {
			panic("No location")
		}

		start := parseTime(propertyMap["DTSTART"])
		end := parseTime(propertyMap["DTEND"])
		if strings.Contains(location, "\\,") {
			for _, roomName := range strings.Split(location, "\\,") {
				appendClass(result, roomName, start, end)
			}
		} else {
			appendClass(result, location, start, end)
		}
	}

	// sort classes by start time
	for _, adeRoom := range result {
		sort.Slice(adeRoom.classes, func(i, j int) bool {
			return adeRoom.classes[i].start.Before(adeRoom.classes[j].start)
		})
	}

	return result
}

func appendClass(result map[string]Room, name string, start, end time.Time) {
	if strings.Contains(name, "Distanciel") {
		return
	}
	name = parseRoomName(name)
	adeRoom, ok := result[name]
	if !ok {
		adeRoom = Room{
			name: name,
		}
	}
	adeRoom.classes = append(adeRoom.classes, Class{
		start: start,
		end:   end,
	})
	result[name] = adeRoom
}

func parseRoomName(name string) string {
	split := strings.Split(name, "ISTIL")
	if strings.Trim(split[0], " ") == "Amphi" {
		return "Amphi"
	} else {
		var number int
		_, err := fmt.Sscan(split[1], &number)
		if err != nil {
			log.Println("Error parsing room name", name)
			return "Unknown"
		}
		return fmt.Sprint(number)
	}
}

func parseTime(value string) time.Time {
	result, err := time.Parse("20060102T150405Z", value)
	if err != nil {
		panic(err)
	}
	return result.Local()
}

func goto20(currentTime time.Time) time.Time {
	timeTo20 := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 20, 0, 0, 0, currentTime.Location()).Sub(currentTime)
	return currentTime.Add(timeTo20)
}

func saveAvailability(from, to time.Time, dbRoom *_ent.Room) {
	if to.Sub(from) > 30*time.Minute {
		ent.Get().AvailableRoom.Create().SetRooms(dbRoom).SetStart(from).SetEnd(to).SaveX(context.Background())
	}
}

func printRooms(rooms map[string]Room) {
	for name, adeRoom := range rooms {
		fmt.Println(name)
		for _, class := range adeRoom.classes {
			fmt.Println("\t", class.start.Format("02/01 \t 15:04 -"), class.end.Format("15:04"))
		}
	}
}

func readableTime(t time.Time) string {
	return t.Format("02/01 15:04")
}

// debug functions
var _ = printRooms
var _ = readableTime
