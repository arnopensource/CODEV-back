package handlers

import (
	"fmt"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/gin-gonic/gin"
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

func GetMyUser(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()
	//userProfile, err := client.Profile.
	//	Query().
	//	Where(profile.ID(user.ID)).
	//	Select(profile.FieldID, profile.FieldFirstname, profile.FieldLastname).
	//	WithFriendsData().
	//	Only(c)
	//if err != nil {
	//	c.AbortWithError(http.StatusInternalServerError, err)
	//	return
	//}

	arno := client.Profile.Query().Where(profile.ID(user.ID)).OnlyX(c)

	bob := client.Profile.Query().Where(profile.ID(uuid.MustParse("9b91d099-2175-42a2-8e43-cb60a7b6182a"))).OnlyX(c)

	friends := arno.QueryFriends().AllX(c)
	fmt.Println(friends)

	friends = bob.QueryFriends().AllX(c)
	fmt.Println(friends)

	//c.JSON(http.StatusOK, userProfile)
	c.JSON(http.StatusOK, "ok")
}
