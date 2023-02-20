package Utility

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"toplearn-api/viewModel/common/security"
)

type ApiContext struct {
	// contain embed of the context
	echo.Context
}

func (c ApiContext) GetUserId() (string, error) {
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*security.JwtClaims)
	return claim.Id, nil
}
