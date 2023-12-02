package middleware

import (
	"github.com/gin-gonic/gin"
	"go-project-frame/pkg/consts"
	"k8s.io/apimachinery/pkg/util/sets"
)

var AlwaysAllowPath sets.String

func InstallMiddleware(ginEngine *gin.RouterGroup) {
	AlwaysAllowPath = sets.NewString(consts.LoginURL, consts.LogoutURL, consts.WebShellURL)
	ginEngine.Use(Logger(), Cors(), Limiter(), Recovery(true), Validator(), JWTAuth(), Casbin())
}
