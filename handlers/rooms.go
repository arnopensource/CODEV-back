package handlers

import (
	"github.com/abc3354/CODEV-back/ent/availableroom"
	"github.com/abc3354/CODEV-back/ent/room"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"

	"github.com/abc3354/CODEV-back/services/ade"
	"github.com/gin-gonic/gin"
)

func GetEmptyRooms(c *gin.Context) {
	at := time.Now()

	atStr := c.Query("time")
	if atStr != "" {
		atInt, err := strconv.Atoi(atStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		at = time.Unix(int64(atInt), 0)
	}

	result := ade.GetEmptyRooms(at)
	c.JSON(http.StatusOK, result)
}

func CreateBooking(c *gin.Context) {
	//auth
	_, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()

	var body BookingBody
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//check if is already booked
	bookedRoom, err := client.Room.Query().Where(room.ID(body.RoomID)).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isAvailable, err := bookedRoom.QueryAvailability().Where(availableroom.And(
		availableroom.StartLTE(body.Start),
		availableroom.EndGTE(body.End),
	)).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if isAvailable == nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": "room is not available",
		})
		return
	}

	//TODO: Check is room is at max capacity

	_, err = client.Booking.Create().
		SetRoomID(body.RoomID).
		SetProfileID(body.ProfileID).
		SetStart(body.Start).
		SetEnd(body.End).
		SetNumber(body.Number).
		Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}
