package controller

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/samarthasthan/lockchain/internal/ethereum"
	"github.com/samarthasthan/lockchain/internal/models"
)

type Controller struct {
	client   *ethclient.Client
	contract *ethereum.Ethereum
}

func NewController(c *ethclient.Client, e *ethereum.Ethereum) *Controller {
	return &Controller{
		client:   c,
		contract: e,
	}
}

// Controllers for handleCreateSecret, handleGetSecrets, and handleGetSecretByIndex

// handleCreateSecret creates a new secret
func (c *Controller) CreateSecret(ctx echo.Context) error {
	var request models.CreateSecretRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	privateKey, err := crypto.HexToECDSA(request.PrivateKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid private key")
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1)) // Replace with appropriate chain ID
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create transaction options")
	}

	tx, err := c.contract.SetSecret(opts, request.Secret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create secret")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"transaction_hash": tx.Hash().Hex(),
	})
}

// handleGetSecrets gets all the secrets
func (c *Controller) GetSecrets(ctx echo.Context) error {
	return nil
}

// handleGetSecretByIndex gets a secret by index
func (c *Controller) GetSecretByIndex(ctx echo.Context) error {
	return nil
}
