// Description: This file contains the handlers for the API endpoints.
// POST /secret
// GET /secrets
// GET /secret/:index

package api

import "github.com/labstack/echo/v4"

func (h *Handlers) handleCreateSecret(c echo.Context) error {
	return h.controller.CreateSecret(c)
}

func (h *Handlers) handleGetSecrets(c echo.Context) error {
	return h.controller.GetSecrets(c)
}

func (h *Handlers) handleGetSecretByIndex(c echo.Context) error {
	return h.controller.GetSecretByIndex(c)
}
