package tests

import (
	"github.com/abc3354/CODEV-back/handlers"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ent.Mock(t)
	defer ent.Close()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handlers.Ping(c)

	if w.Code != 200 {
		t.Error(w.Code)
	}

	if w.Body.String() != "\"pong\"" {
		t.Error(w.Body.String())
	}
}

func TestPingEnt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ent.Mock(t)
	defer ent.Close()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handlers.PingEnt(c)

	if w.Code != 200 {
		t.Error(w.Code)
	}

	if w.Body.String() != "\"ent ok\"" {
		t.Error(w.Body.String())
	}
}
