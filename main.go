package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"
)

var (
	supabaseProjectRef string
	supabaseAPIKey     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	supabaseProjectRef = os.Getenv("SUPABASE_PROJECT_REF")
	supabaseAPIKey = os.Getenv("SUPABASE_API_KEY")

	if supabaseProjectRef == "" || supabaseAPIKey == "" {
		log.Fatal(".env file is incomplete")
	}
}

func main() {
	authClient := gotrue.New(supabaseProjectRef, supabaseAPIKey)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/login", func(c *gin.Context) {
		var body LoginBody
		c.MustBindWith(&body, binding.JSON)
		resp, err := authClient.SignInWithEmailPassword(body.Email, body.Password)
		if err != nil {
			// todo check if it is a login error
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, resp.AccessToken)
	})
	r.POST("/signup", func(c *gin.Context) {
		var body SignupBody
		c.MustBindWith(&body, binding.JSON)
		resp, err := authClient.Signup(types.SignupRequest{
			Email:    body.Email,
			Password: body.Password,
			Data:     map[string]interface{}{"nfc": body.NFC},
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, resp.AccessToken)
	})

	r.Run()
}

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	NFC      string `json:"nfc" binding:"required"`
}
