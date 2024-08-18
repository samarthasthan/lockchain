package models

import (
	"time"
)

type SecretStruct struct {
	ID        int
	Secret    string
	CreatedAt time.Time
}
