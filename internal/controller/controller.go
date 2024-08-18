package controller

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
// CreateSecret creates a new secret in the contract
func (c *Controller) CreateSecret(ctx echo.Context) error {
	var request models.CreateSecretRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Ensure private key is in the correct format
	privateKey, err := crypto.HexToECDSA(request.PrivateKey[2:]) // Remove '0x' prefix
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid private key")
	}

	// Create transaction options
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Replace with the correct chain ID
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create transaction options")
	}

	// Interact with the smart contract
	tx, err := c.contract.SetSecret(opts, request.Secret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create secret")
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"transaction_hash": tx.Hash().Hex(),
	})
}

// handleGetSecrets gets all the secrets
// GetSecrets retrieves all secrets from the contract
// GetSecrets retrieves all secrets for a given user address from the contract
func (c *Controller) GetSecrets(ctx echo.Context) error {
	var req models.GetSecretsRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Ensure the address is in the correct format
	address := req.Address
	if !common.IsHexAddress(address) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid address format")
	}

	addr := common.HexToAddress(address)

	// Retrieve secrets from the contract
	secrets, err := c.contract.GetSecrets(&bind.CallOpts{From: addr})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve secrets")
	}

	// Return the secrets
	return ctx.JSON(http.StatusOK, secrets)
}

// handleGetSecretByIndex gets a secret by index
// GetSecretByIndex retrieves a specific secret by index for a given user address from the contract
func (c *Controller) GetSecretByIndex(ctx echo.Context) error {
	// Index from path parameter
	index := ctx.Param("index")
	var req models.GetSecretsRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	// Ensure the address is in the correct format
	address := req.Address
	if !common.IsHexAddress(address) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid address format")
	}

	addr := common.HexToAddress(address)

	// Validate index
	if index == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Index is required")
	}

	// Big int from index
	i, err := strconv.Atoi(index)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid index")
	}
	// Retrieve the secret by index from the contract
	secret, err := c.contract.GetSecretByIndex(&bind.CallOpts{From: addr}, big.NewInt(int64(i)))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve secret")
	}

	// Return the secret
	return ctx.JSON(http.StatusOK, secret)
}
