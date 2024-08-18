package models

import (
	"math/big"
)

// Payload for creating a secret
type CreateSecretRequest struct {
	Address    string `json:"address" validate:"required"`
	PrivateKey string `json:"private_key" validate:"required"`
	Secret     string `json:"secret" validate:"required"`
}

// Payload for getting secrets
type GetSecretsRequest struct {
	Address string `json:"address" validate:"required"`
}

// Payload for getting a secret by index
type GetSecretByIndexRequest struct {
	Address string   `json:"address" validate:"required"`
	Index   *big.Int `json:"index" validate:"required"`
}
