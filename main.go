package main

import (
	"log"

	"github.com/abc3354/CODEV-back/handlers"
	"github.com/abc3354/CODEV-back/services/ade"
	"github.com/abc3354/CODEV-back/services/database"
	"github.com/abc3354/CODEV-back/services/ent"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	defer database.Close()
	defer ent.Close()
	ade.Update()

	r := gin.Default()

	r.GET("/", handlers.Home)
	r.GET("/ping", handlers.Ping)
	r.GET("/ping/ent", handlers.PingEnt)

	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/nfc/login", handlers.NFCLogin)
	r.PUT("/auth/nfc/id", handlers.NFCChangeID)
	r.GET("/auth/check", handlers.CheckToken)

	r.GET("/rooms/empty", handlers.GetEmptyRooms)
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserById)
	r.POST("/bookings", handlers.CreateBooking)
	r.PUT("/users/:id", handlers.UptadeUsers)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
