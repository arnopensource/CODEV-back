package main

import (
	"fmt"
	ics "github.com/arran4/golang-ical"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
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

var db = make(map[string]Room)

func updateDB() {
	downloadURL := generateURL(time.Now())

	client := http.Client{
		Timeout: 5 * time.Minute,
	}
	res, err := client.Get(downloadURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	db = parseICAL(res.Body)
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

		roomName, ok := propertyMap["LOCATION"]
		if !ok {
			panic("No location")
		}

		if strings.Contains(roomName, "\\,") {
			for _, roomName := range strings.Split(roomName, "\\,") {
				if strings.Contains(roomName, "Distanciel") {
					continue
				}
				r, ok := result[roomName]
				if !ok {
					r = Room{
						name: roomName,
					}
				}
				r.classes = append(r.classes, Class{
					start: parseTime(propertyMap["DTSTART"]),
					end:   parseTime(propertyMap["DTEND"]),
				})
				result[roomName] = r
			}
		} else {
			r, ok := result[roomName]
			if !ok {
				r = Room{
					name: roomName,
				}
			}
			r.classes = append(r.classes, Class{
				start: parseTime(propertyMap["DTSTART"]),
				end:   parseTime(propertyMap["DTEND"]),
			})
			result[roomName] = r
		}
	}

	for _, room := range result {
		sort.Slice(room.classes, func(i, j int) bool {
			return room.classes[i].start.Before(room.classes[j].start)
		})
	}

	return result
}

func printRooms() {
	for name, room := range db {
		fmt.Println(name)
		for _, class := range room.classes {
			fmt.Println("\t", class.start.Format("02/01 \t 15:04 -"), class.end.Format("15:04"))
		}
	}
}

func parseTime(value string) time.Time {
	result, err := time.Parse("20060102T150405Z", value)
	if err != nil {
		panic(err)
	}
	return result.Local()
}

type EmptyRoom struct {
	Name     string `json:"name"`
	FreeTime int    `json:"freeTime"`
}

func getEmptyRooms(at time.Time) []EmptyRoom {
	result := make([]EmptyRoom, 0)
	for _, room := range db {
		var closest time.Time
		for _, class := range room.classes {
			if class.end.Before(at) {
				continue
			}
			if closest.IsZero() || class.start.Before(closest) {
				closest = class.start
			}
		}

		if closest.Sub(at) > time.Hour {
			result = append(result, EmptyRoom{
				Name:     room.name,
				FreeTime: int(closest.Sub(at).Minutes()),
			})
		}
	}
	return result
}
