package handlers

import (
	"net/http"
	"time"

	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/services/ent"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
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
		Where(profile.Or(
			profile.FirstnameContainsFold(filterName),
			profile.LastnameContainsFold(filterName),
		)).
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

func GetMyUser(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()
	userProfile, err := client.Profile.
		Query().
		Where(profile.ID(user.ID)).
		Select(profile.FieldID, profile.FieldFirstname, profile.FieldLastname).
		Only(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// add friends manually because there is a problem with .WithFriends
	for _, req := range getFriendsByUserId(user.ID) {
		friends, err := req.All(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		userProfile.Edges.Friends = append(userProfile.Edges.Friends, friends...)
	}

	c.JSON(http.StatusOK, userProfile)
}

func UpdateUser(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var body UpdateUserBody
	if err = c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()
	err = client.Profile.
		UpdateOneID(user.ID).
		SetFirstname(body.Firstname).
		SetLastname(body.Lastname).
		Exec(c)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func AddFriend(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	friendID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()
	_, err = client.Friend.
		Create().
		SetProfileID(user.ID).
		SetFriendID(friendID).
		SetAccepted(false).
		SetSince(time.Now()).
		Save(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusCreated)
}
