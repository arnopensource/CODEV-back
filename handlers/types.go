package handlers

import "time"

type LoginBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	NFC      string `json:"nfc" binding:"required"`
}

type SignupResponse struct {
	Email  string     `json:"email"`
	SentAt *time.Time `json:"sent_at"`
}

type NFCLoginBody struct {
	NFC string `json:"nfc" binding:"required"`
}

type NFCModificationBody struct {
	Token string `json:"token" binding:"required"`
	NFC   string `json:"nfc" binding:"required"`
}
