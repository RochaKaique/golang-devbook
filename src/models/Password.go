package models

type Password struct {
	NewPassword    string `json:"nova,omitempty"`
	ActualPassword string `json:"atual,omitempty"`
}
