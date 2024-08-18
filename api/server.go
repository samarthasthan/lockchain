package api

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/samarthasthan/lockchain/internal/controller"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
	Handlers struct {
		*echo.Echo
		controller *controller.Controller
	}
)

// Validate validates the input
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// NewHandler creates a new handler
func NewHandler(c *controller.Controller) *Handlers {
	return &Handlers{Echo: echo.New(), controller: c}
}

// RegisterValidator registers the custom validator
func (h *Handlers) RegisterValidator() {
	h.Validator = &CustomValidator{validator: validator.New()}
}

// Handle handles the routes
func (h *Handlers) Handle() {
	// Handle the authentication routes
	h.RegisterValidator()
	h.POST("/secret", h.handleCreateSecret)
	h.GET("/secrets", h.handleGetSecrets)
	h.GET("/secret/:index", h.handleGetSecretByIndex)
}
