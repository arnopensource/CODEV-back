package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/abc3354/CODEV-back/ent/profile"
	"github.com/abc3354/CODEV-back/services/database"
	"github.com/abc3354/CODEV-back/services/ent"
	"github.com/abc3354/CODEV-back/services/env"
	"github.com/abc3354/CODEV-back/services/gotrue"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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
	if err = createProfile(c, resp.User.ID); err != nil {
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
	if err = createProfile(c, resp.User.ID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
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
	var userID uuid.UUID
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
	if err = createProfile(c, userID); err != nil {
		return
	}
	c.JSON(http.StatusOK, token)
}

func NFCChangeID(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": "no token in header",
		})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")

	var body NFCModificationBody
	if err := c.MustBindWith(&body, binding.JSON); err != nil {
		return
	}
	_, err := gotrue.Get().WithToken(token).UpdateUser(types.UpdateUserRequest{
		Data: map[string]interface{}{"nfc": body.NFC},
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func CheckToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if !strings.HasPrefix(token, "Bearer ") {
		c.JSON(http.StatusBadRequest, map[string]any{
			"error": "no token in header",
		})
		return
	}
	token = strings.TrimPrefix(token, "Bearer ")

	_, err := gotrue.Get().WithToken(token).GetUser()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func createProfile(c *gin.Context, id uuid.UUID) error {
	client := ent.Get()
	hasProfile, err := client.Profile.Query().Where(profile.ID(id)).Count(c)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return err
	}
	if hasProfile == 0 {
		_, err = client.Profile.Create().SetID(id).SetFirstname("").SetLastname("").Save(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return err
		}
	}
	return nil
}
