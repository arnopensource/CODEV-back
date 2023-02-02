package handlers

import (
	"context"
	_ent "github.com/abc3354/CODEV-back/ent"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"

	"github.com/abc3354/CODEV-back/ent/availableroom"
	"github.com/abc3354/CODEV-back/ent/event"
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

	isAvailable, err := checkRoom(body.RoomID, body.Start, body.End, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !isAvailable {
		c.JSON(http.StatusConflict, gin.H{"message": "Room is not available"})
		return
	}

	creatorProfile, err := client.Profile.Query().Where(profile.ID(user.ID)).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	newEvent, err := client.Event.Create().
		SetName(body.Name).
		SetActivity(body.Activity).
		SetStart(body.Start).
		SetEnd(body.End).
		SetRoomID(body.RoomID).
		AddProfiles(creatorProfile).
		Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	firstMember, err := client.Member.Query().Where(member.And(member.EventIDIn(newEvent.ID), member.ProfileID(creatorProfile.ID))).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err = firstMember.Update().SetIsAdmin(true).Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, newEvent)
}

func checkRoom(roomID int, start time.Time, end time.Time, ctx context.Context) (bool, error) {
	client := ent.Get()

	bookedRoom, err := client.Room.Query().Where(room.ID(roomID)).Only(ctx)
	if err != nil {
		return false, err
	}

	isAvailable, err := bookedRoom.QueryAvailability().Where(availableroom.And(
		availableroom.StartLTE(start),
		availableroom.EndGTE(end),
	)).Only(ctx)
	if err != nil {
		return false, err
	}

	if isAvailable == nil {
		return false, nil
	}

	return true, nil
}

func GetMyEvents(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()

	userProfile, err := client.Profile.Query().Where(profile.ID(user.ID)).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	events, err := userProfile.QueryEvents().WithMembers(func(query *_ent.MemberQuery) {
		query.Where(member.ProfileID(user.ID))
	}).All(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var result []struct {
		*_ent.Event
		IsAdmin bool `json:"isAdmin"`
	}
	for _, extractedEvent := range events {
		isAdmin := extractedEvent.Edges.Members[0].IsAdmin
		extractedEvent.Edges.Members = nil
		result = append(result, struct {
			*_ent.Event
			IsAdmin bool `json:"isAdmin"`
		}{extractedEvent, isAdmin})
	}

	c.JSON(http.StatusOK, result)
}

type ModifyEventRequest struct {
	Activity string    `json:"activity"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	RoomID   int       `json:"roomId"`
}

func ModifyEvent(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var body ModifyEventRequest
	if err = c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()

	isAdmin, err := checkAdminRights(user.ID, eventID, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !isAdmin {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"error": "you are not admin",
		})
		return
	}

	eventToModify, err := client.Event.Query().Where(event.ID(eventID)).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	update := eventToModify.Update()

	if body.Activity != "" {
		update = update.SetActivity(body.Activity)
	}

	if !body.Start.IsZero() {
		update = update.SetStart(body.Start)
	}

	if !body.End.IsZero() {
		update = update.SetEnd(body.End)
	}

	if body.RoomID != 0 {
		isAvailable, err := checkRoom(body.RoomID, body.Start, body.End, c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !isAvailable {
			c.JSON(http.StatusConflict, gin.H{"message": "Room is not available"})
			return
		}
		update = update.SetRoomID(body.RoomID)
	}

	modifiedEvent, err := update.Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, modifiedEvent)
}

func checkAdminRights(userId uuid.UUID, eventId int, ctx context.Context) (bool, error) {
	client := ent.Get()

	userMembership, err := client.Member.Query().Where(member.And(member.EventID(eventId), member.ProfileID(userId))).Only(ctx)
	if err != nil {
		return false, err
	}

	return userMembership.IsAdmin, nil
}

func CancelEvent(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	eventID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()

	isAdmin, err := checkAdminRights(user.ID, eventID, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !isAdmin {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"error": "you are not admin",
		})
		return
	}

	_, err = client.Member.Delete().Where(member.EventID(eventID)).Exec(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	_, err = client.Event.Delete().Where(event.ID(eventID)).Exec(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "Event deleted",
	})
}
