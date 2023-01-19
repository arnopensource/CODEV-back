package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/supabase-community/gotrue-go"
	"github.com/supabase-community/gotrue-go/types"

	_ "github.com/lib/pq"
)

var (
	supabaseProjectRef string
	supabaseAPIKey     string
	datasource         string
	jwtSecret          string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	supabaseProjectRef = os.Getenv("SUPABASE_PROJECT_REF")
	supabaseAPIKey = os.Getenv("SUPABASE_API_KEY")
	datasource = os.Getenv("DATASOURCE")
	jwtSecret = os.Getenv("JWT_SECRET")

	if supabaseProjectRef == "" || supabaseAPIKey == "" || datasource == "" || jwtSecret == "" {
		log.Fatal("Environment is incomplete")
	}
}

func main() {
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	authClient := gotrue.New(supabaseProjectRef, supabaseAPIKey)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
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
	r.POST("/nfc/login", func(c *gin.Context) {
		var body NFCLoginBody
		c.MustBindWith(&body, binding.JSON)
		var userID string
		err = db.QueryRowContext(context.Background(), "SELECT id FROM auth.users WHERE raw_user_meta_data->>'nfc' = $1", body.NFC).Scan(&userID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"aud": "authenticated",
			"sub": userID,
		}).SignedString([]byte(jwtSecret))
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, token)
	})
	r.PUT("/nfc/id", func(c *gin.Context) {
		var body NFCModificationBody
		c.MustBindWith(&body, binding.JSON)

		_, err := authClient.WithToken(body.Token).UpdateUser(types.UpdateUserRequest{
			Data: map[string]interface{}{"nfc": body.NFC},
		})
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, "ok")
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

type NFCLoginBody struct {
	NFC string `json:"nfc" binding:"required"`
}

type NFCModificationBody struct {
	Token string `json:"token" binding:"required"`
	NFC   string `json:"nfc" binding:"required"`
}
