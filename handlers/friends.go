package handlers

import (
	"net/http"
	"time"

	_ent "github.com/abc3354/CODEV-back/ent"
	"github.com/abc3354/CODEV-back/ent/friend"
	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func getFriendsByUserId(id uuid.UUID) []*_ent.ProfileQuery {
	client := ent.Get()
	req := client.Friend.
		Query().
		Where(friend.And(
			friend.Or(
				friend.ProfileID(id),
				friend.FriendID(id),
			),
			friend.Accepted(true),
		))

	return []*_ent.ProfileQuery{req.QueryFriend().Where(profile.IDNEQ(id)), req.QueryProfile().Where(profile.IDNEQ(id))}
}

func AddFriend(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var body FriendRequestBody
	if err = c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	friendID, err := uuid.Parse(body.ID)
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

func GetFriends(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	profiles := []*_ent.Profile{}

	for _, req := range getFriendsByUserId(user.ID) {
		friends, err := req.All(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		profiles = append(profiles, friends...)
	}

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, profiles)
}

func FriendRequestDecision(c *gin.Context) {
	var body FriendRequestDecisionBody

	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()

	friendID, err := uuid.Parse(body.ID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	friendPredicate :=
		friend.And(
			friend.ProfileID(friendID),
			friend.FriendID(user.ID),
		)

	if *body.Accepted {
		friendRequest, err := client.Friend.
			Query().
			Where(friendPredicate).
			Only(c)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// If friend request already accepted
		if friendRequest.Accepted {
			c.AbortWithStatus(http.StatusConflict)
			return
		}

		_, err = client.Friend.
			UpdateOne(friendRequest).
			SetAccepted(true).
			SetSince(time.Now()).
			Save(c)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, true)
	} else {
		_, err = client.Friend.
			Delete().Where(friendPredicate).
			Exec(c)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, false)
	}
}

func RemoveFriend(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	var body FriendRequestBody
	if err = c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	friendID, err := uuid.Parse(body.ID)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	client := ent.Get()

	friendPredicate :=
		friend.And(
			friend.Or(
				friend.ProfileID(user.ID),
				friend.FriendID(user.ID),
			),
			friend.Or(
				friend.ProfileID(friendID),
				friend.FriendID(friendID),
			),
		)

	_, err = client.Friend.
		Delete().Where(friendPredicate).
		Exec(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func GetFriendRequests(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()

	req, err := client.Friend.
		Query().
		Where(friend.And(
			friend.FriendID(user.ID),
			friend.Accepted(false),
		)).
		WithProfile().
		All(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, req)
}

func GetSentFriendRequests(c *gin.Context) {
	user, err := checkToken(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := ent.Get()

	req, err := client.Friend.
		Query().
		Where(friend.And(
			friend.ProfileID(user.ID),
			friend.Accepted(false),
		)).
		WithFriend().
		All(c)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, req)
}
