package main

import (
	"log"

	"github.com/abc3354/CODEV-back/handlers"
	"github.com/abc3354/CODEV-back/services/ade"
	"github.com/abc3354/CODEV-back/services/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	defer database.Close()
	ade.Update()

	r := gin.Default()

	r.GET("/ping", handlers.Ping)

	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/nfc/login", handlers.NFCLogin)
	r.PUT("/auth/nfc/id", handlers.NFCChangeID)

	r.GET("/rooms/empty", handlers.GetEmptyRooms)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
