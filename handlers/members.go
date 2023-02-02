package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/abc3354/CODEV-back/ent/event"
	"github.com/abc3354/CODEV-back/services/ent"
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
