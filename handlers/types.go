package handlers

import (
	"github.com/google/uuid"
	"time"
)

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	NFC      string `json:"nfc"`
}

type SignupResponse struct {
	Email  string     `json:"email"`
	SentAt *time.Time `json:"sent_at"`
}

type NFCLoginBody struct {
	NFC string `json:"nfc" binding:"required"`
}

type NFCModificationBody struct {
	NFC string `json:"nfc" binding:"required"`
}

type BookingBody struct {
	ProfileID uuid.UUID `json:"profile_id"`
	RoomID    int       `json:"room_id"`
	Number    int       `json:"number"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}
