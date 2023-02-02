package tests

import (
	"github.com/abc3354/CODEV-back/handlers"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http/httptest"
	"testing"
)

func TestGetUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ent.Mock(t)
	defer ent.Close()

	userID := uuid.New()

	ent.Get().Profile.Create().
		SetID(userID).
		SetEmail("test@example.org")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{
		gin.Param{Key: "id", Value: userID.String()},
	}

	// todo trouver comment mettre une authorization dans le header

	handlers.GetUserById(c)

	if w.Code != 200 {
		t.Error(w.Code)
	}

	log.Println(w.Body.String())
}
