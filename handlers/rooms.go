package handlers

import (
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
	client := ent.Get()

	var body BookingBody
	if err := c.MustBindWith(&body, binding.JSON); err != nil {
		return
	}

	_, err := client.Booking.Create().
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
