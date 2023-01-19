package main

import (
	"github.com/abc3354/CODEV-back/services/ade"
	"log"
	"time"

	"github.com/abc3354/CODEV-back/handlers"
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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/ping", handlers.Ping)

	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/nfc/login", handlers.NFCLogin)
	r.PUT("/auth/nfc/id", handlers.NFCChangeID)

	r.GET("/rooms/empty", handlers.GetEmptyRooms)

	r.GET("/test", func(c *gin.Context) {
		booking, err := ent.Get().Booking.
			Create().
			SetName("204").
			SetStart(time.Now()).
			SetEnd(time.Now().Add(time.Hour)).
			Save(c)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, booking)
	})

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
