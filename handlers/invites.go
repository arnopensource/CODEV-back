package handlers

import (
	"context"
	"net/http"
	"strconv"

	_ent "github.com/abc3354/CODEV-back/ent"
	"github.com/abc3354/CODEV-back/ent/eventinvite"
	"github.com/abc3354/CODEV-back/ent/member"
	"github.com/abc3354/CODEV-back/services/ent"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SendInviteRequest struct {
	EventID   int       `json:"eventId" binding:"required"`
	ProfileID uuid.UUID `json:"userId" binding:"required"`
}

func SendInvite(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var body SendInviteRequest
	if err = c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isAdmin, err := checkAdminRights(user.ID, body.EventID, c)
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

	client := ent.Get()

	exists, err := checkInvite(body.EventID, body.ProfileID, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if exists {
		c.JSON(http.StatusConflict, map[string]any{
			"error": "invite already exists",
		})
		return
	}

	exists, err = checkMember(body.EventID, body.ProfileID, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if exists {
		c.JSON(http.StatusConflict, map[string]any{
			"error": "user is already a member",
		})
		return
	}

	sentInvite, err := client.EventInvite.Create().SetEventID(body.EventID).SetProfileID(body.ProfileID).Save(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, sentInvite)
}

func checkInvite(eventId int, userId uuid.UUID, c context.Context) (bool, error) {
	client := ent.Get()

	_, err := client.EventInvite.Query().Where(eventinvite.EventID(eventId), eventinvite.ProfileID(userId)).Only(c)
	if err == nil {
		return true, nil
	} else if _, ok := err.(*_ent.NotFoundError); ok {
		return false, nil
	} else {
		return false, err
	}
}

func checkMember(eventId int, userId uuid.UUID, c context.Context) (bool, error) {
	client := ent.Get()

	_, err := client.Member.Query().Where(member.EventID(eventId), member.ProfileID(userId)).Only(c)
	if err == nil {
		return true, nil
	} else if _, ok := err.(*_ent.NotFoundError); ok {
		return false, nil
	} else {
		return false, err
	}
}

func GetMyInvites(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()

	invites, err := client.EventInvite.Query().Where(eventinvite.ProfileID(user.ID)).WithEvent().All(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, invites)
}

type InviteDecisionRequest struct {
	Accept *bool `json:"accept" binding:"required"`
}

func InviteDecision(c *gin.Context) {
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

	var body InviteDecisionRequest
	if err = c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()

	exists, err := checkInvite(eventID, user.ID, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !exists {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	_, err = client.EventInvite.Delete().Where(eventinvite.EventID(eventID), eventinvite.ProfileID(user.ID)).Exec(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if *body.Accept {
		_, err = client.Member.Create().SetEventID(eventID).SetProfileID(user.ID).Save(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

type CancelInviteRequest struct {
	EventID   int       `json:"eventId" binding:"required"`
	ProfileID uuid.UUID `json:"profileId" binding:"required"`
}

func CancelInvite(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var body CancelInviteRequest
	if err = c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isAdmin, err := checkAdminRights(user.ID, body.EventID, c)
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

	exists, err := checkInvite(body.EventID, body.ProfileID, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !exists {
		c.JSON(http.StatusConflict, map[string]any{
			"error": "invite does not exist",
		})
		return
	}

	client := ent.Get()

	_, err = client.EventInvite.Delete().Where(eventinvite.EventID(body.EventID), eventinvite.ProfileID(body.ProfileID)).Exec(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func GetInvitesOfEvent(c *gin.Context) {
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

	client := ent.Get()

	invites, err := client.EventInvite.Query().Where(eventinvite.EventID(eventID)).WithProfile().All(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, invites)
}
