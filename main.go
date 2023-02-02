package main

import (
	"log"

	"github.com/abc3354/CODEV-back/handlers"
	"github.com/abc3354/CODEV-back/services/database"
	"github.com/abc3354/CODEV-back/services/ent"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	defer database.Close()
	defer ent.Close()

	r := gin.Default()

	r.GET("/", handlers.Home)
	r.GET("/ping", handlers.Ping)
	r.GET("/ping/ent", handlers.PingEnt)

	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/nfc/login", handlers.NFCLogin)
	r.PUT("/auth/nfc/id", handlers.NFCChangeID)
	r.GET("/auth/check", handlers.CheckToken)

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserById)
	r.GET("/users/me", handlers.GetMyUser)
	r.PUT("/users", handlers.UpdateUser)

	r.GET("/rooms/empty", handlers.GetEmptyRooms)

	r.POST("/bookings", handlers.CreateBooking)
	r.GET("/bookings", handlers.GetUserBookings)

	r.GET("/friends", handlers.GetFriends)
	r.POST("/friends", handlers.AddFriend)
	r.PUT("/friends", handlers.FriendRequestDecision)
	r.DELETE("/friends", handlers.RemoveFriend)
	r.GET("/friends/requests", handlers.GetFriendRequests)
	r.GET("/friends/requests/sent", handlers.GetSentFriendRequests)

	r.POST("/events", handlers.CreateEvent)
	r.GET("/events", handlers.GetMyEvents)
	r.PUT("/events/:id", handlers.ModifyEvent)
	r.DELETE("/events/:id", handlers.CancelEvent)

	r.POST("/events/invites", handlers.SendInvite)
	r.GET("/events/invites", handlers.GetMyInvites)
	r.PUT("/events/invites/:id", handlers.InviteDecision)
	r.DELETE("/events/invites", handlers.CancelInvite)
	r.GET("/events/{id}/invites", nil)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
