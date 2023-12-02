package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/consts"
	"go-project-frame/pkg/globalError"
	"go-project-frame/pkg/jwt"
)

func GetClaims(ctx *gin.Context) (*jwt.CustomClaims, error) {
	token := ctx.Request.Header.Get(consts.TokenKey)
	if token == "" {
		return nil, globalError.NewGlobalError(globalError.AuthorizationLackToken, errors.New(globalError.GetCodeTag(globalError.AuthorizationLackToken)))
	}
	claims, err := jwt.JwtToken.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
