package model

import "time"

type Credentials struct {
	User           string
	Token          string
	ExpirationTime time.Time
	Role           string
}
