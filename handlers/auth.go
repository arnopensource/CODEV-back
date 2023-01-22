package handlers

import (
	"context"
	"github.com/abc3354/CODEV-back/services/database"
	"github.com/abc3354/CODEV-back/services/env"
	"github.com/abc3354/CODEV-back/services/gotrue"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"github.com/supabase-community/gotrue-go/types"
)

func Login(c *gin.Context) {
	var body LoginBody
	if err := c.MustBindWith(&body, binding.JSON); err != nil {
		return
	}
	resp, err := gotrue.Get().SignInWithEmailPassword(body.Email, body.Password)
	if err != nil {
		// todo check if it is a login error
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, resp.AccessToken)
}

func Signup(c *gin.Context) {
	var body SignupBody
	if err := c.MustBindWith(&body, binding.JSON); err != nil {
		return
	}
	req := types.SignupRequest{
		Email:    body.Email,
		Password: body.Password,
	}
	if body.NFC != "" {
		req.Data = map[string]interface{}{"nfc": body.NFC}
	}
	resp, err := gotrue.Get().Signup(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if !resp.ConfirmedAt.IsZero() {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": "email already in use",
		})
		return
	}
	c.JSON(http.StatusOK, SignupResponse{
		Email:  resp.Email,
		SentAt: resp.ConfirmationSentAt,
	})
}

func NFCLogin(c *gin.Context) {
	var body NFCLoginBody
	if err := c.MustBindWith(&body, binding.JSON); err != nil {
		return
	}
	var userID string
	err := database.Get().QueryRowContext(context.Background(), "SELECT id FROM auth.users WHERE raw_user_meta_data->>'nfc' = $1", body.NFC).Scan(&userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": "authenticated",
		"sub": userID,
	}).SignedString([]byte(env.Get().JWTSecret))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func NFCChangeID(c *gin.Context) {
	var body NFCModificationBody
	if err := c.MustBindWith(&body, binding.JSON); err != nil {
		return
	}
	_, err := gotrue.Get().WithToken(body.Token).UpdateUser(types.UpdateUserRequest{
		Data: map[string]interface{}{"nfc": body.NFC},
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func CheckToken(c *gin.Context) {
	var body TokenBody
	if err := c.MustBindWith(&body, binding.JSON); err != nil {
		return
	}
	_, err := gotrue.Get().WithToken(body.Token).GetUser()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
