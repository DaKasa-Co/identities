package model

import (
	"github.com/google/uuid"
)

// Identity represents the global scope request about identities service
type Identity struct {
	ID           uuid.UUID    `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	Username     string       `json:"username,omitempty"`
	Email        string       `json:"email,omitempty"`
	Password     string       `json:"password,omitempty"`
	Address      string       `json:"address,omitempty"`
	PhoneNumber  int          `json:"phoneNumber,omitempty"`
	Birthday     string       `json:"birthday,omitempty"`
	Avatar       string       `json:"avatar,omitempty"`
	SocialMedias SocialMedias `json:"socialMedias,omitempty"`
	Stamp        Stamp        `json:"stamp,omitempty"`
	CreatedAt    string       `json:"createdAt,omitempty"`
	UpdatedAt    string       `json:"updatedAt,omitempty"`
	Status       Status       `json:"status,omitempty"`
}

// Stamp represents DaKasa recognition stamps
type Stamp struct {
	IsVerified bool   `json:"isverified,omitempty"`
	Type       string `json:"type,omitempty"`
}

// Status represents request status
type Status struct {
	Ticket     uuid.UUID        `json:"ticket,omitempty"`
	Validation StatusValidation `json:"validation,omitempty"`
}

// StatusValidation represents the validation fields
type StatusValidation struct {
	Tmp      int    `json:"tmp,omitempty"`
	Password string `json:"password,omitempty"`
}

// SocialMedias represents the User's Social Medias
type SocialMedias struct {
	Instagram string `json:"instagram,omitempty"`
	Facebook  string `json:"facebook,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	TikTok    string `json:"tiktok,omitempty"`
	Kwai      string `json:"kwai,omitempty"`
	LinkedIn  string `json:"linkedIn,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Twitch    string `json:"twitch,omitempty"`
}
