package handlers

import (
	"net/http"
	"strconv"

	"github.com/abc3354/CODEV-back/ent/event"
	"github.com/abc3354/CODEV-back/ent/member"
	"github.com/abc3354/CODEV-back/services/ent"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetEventMembers(c *gin.Context) {
	_, err := checkToken(c)
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

	targetEvent, err := client.Event.Query().Where(event.ID(eventID)).Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	members, err := client.Event.QueryMembers(targetEvent).WithProfile().All(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, members)
}

type Member struct {
	ProfileID uuid.UUID `json:"profileId"`
}

func RemoveMember(c *gin.Context) {
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

	var body Member
	if err = c.ShouldBindJSON(&body); err != nil {
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

	isAdmin, err = checkAdminRights(body.ProfileID, eventID, c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if isAdmin {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"error": "you can't remove an admin",
		})
		return
	}

	client := ent.Get()

	_, err = client.Member.Delete().Where(member.ProfileID(body.ProfileID), member.EventID(eventID)).Exec(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}

func QuitEvent(c *gin.Context) {
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
	if isAdmin {
		adminNumber, err := client.Member.Query().Where(member.And(member.EventID(eventID), member.IsAdmin(true))).Count(c)
		if err != nil {
			return
		}
		if adminNumber == 1 {
			c.JSON(http.StatusUnauthorized, map[string]any{
				"error": "you are the last admin",
			})
			return
		}
	}

	_, err = client.Member.Delete().Where(member.ProfileID(user.ID), member.EventID(eventID)).Exec(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"message": "success",
	})
}
