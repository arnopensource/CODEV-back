package handlers

import (
	"net/http"

	"github.com/abc3354/CODEV-back/services/ent"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

// PingEnt will do ent migration if ent was not used yet
func PingEnt(_ *gin.Context) {
	ent.Get()
}

func Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello !")
}
