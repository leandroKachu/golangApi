package model

type Password struct {
	NewPassword     string `json:"newpassword"`
	CurrentPassword string `json:"currentpassword"`
}
