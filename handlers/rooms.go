package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/abc3354/CODEV-back/services/ade"
	"github.com/gin-gonic/gin"
)

func GetEmptyRooms(c *gin.Context) {
	at := time.Now()

	atStr := c.Query("time")
	if atStr != "" {
		atInt, err := strconv.Atoi(atStr)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		at = time.Unix(int64(atInt), 0)
	}

	result := ade.GetEmptyRooms(at)
	c.JSON(http.StatusOK, result)
}
