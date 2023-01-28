package handlers

import (
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func GetUsers(c *gin.Context) {

	client := ent.Get()
	profiles, err := client.Profile.
		Query().
		Select(profile.FieldID, profile.FieldFirstname, profile.FieldLastname).
		All(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, profiles)

}

func GetUserById(c *gin.Context) {

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()
	userProfile, err := client.Profile.
		Query().
		Where(profile.ID(userID)).
		Select(profile.FieldID, profile.FieldFirstname, profile.FieldLastname).
		Only(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, userProfile)

}
