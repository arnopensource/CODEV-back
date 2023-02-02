package handlers

import (
	"net/http"
	"time"

	"github.com/abc3354/CODEV-back/ent/availableroom"
	"github.com/abc3354/CODEV-back/ent/member"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/ent/room"
	"github.com/abc3354/CODEV-back/services/ent"

	"github.com/gin-gonic/gin"
)

type CreateEventRequest struct {
	Name     string    `json:"name" binding:"required"`
	Activity string    `json:"activity" binding:"required"`
	Start    time.Time `json:"start" binding:"required"`
	End      time.Time `json:"end" binding:"required"`
	RoomID   int       `json:"roomId" binding:"required"`
}

func CreateEvent(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var body CreateEventRequest
	if err = c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()

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
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if isAvailable == nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": "room is not available",
		})
		return
	}

	creatorProfile, err := client.Profile.Query().Where(profile.ID(user.ID)).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	event, err := client.Event.Create().
		SetName(body.Name).
		SetActivity(body.Activity).
		SetStart(body.Start).
		SetEnd(body.End).
		SetRoom(bookedRoom).
		AddProfiles(creatorProfile).
		Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	firstMember, err := client.Member.Query().Where(member.And(member.EventIDIn(event.ID), member.ProfileID(creatorProfile.ID))).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err = firstMember.Update().SetIsAdmin(true).Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, event)
}
