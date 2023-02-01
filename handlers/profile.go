package handlers

import (
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"net/http"
)

func GetUsers(c *gin.Context) {
	_, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	filterName := c.Query("name")

	client := ent.Get()
	profiles, err := client.Profile.
		Query().
		Where(profile.FirstnameContains(filterName)).
		Select(profile.FieldID, profile.FieldFirstname, profile.FieldLastname).
		All(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, profiles)

}

func GetUserById(c *gin.Context) {

	_, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

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

func UptadeUsers(c *gin.Context) {
	_, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var body UpdateUserBody
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()
	err = client.Profile.
		UpdateOneID(userID).
		SetFirstname(body.Firstname).
		SetLastname(body.Lastname).
		Exec(c)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
