package controller

import (
	"github.com/gin-gonic/gin"
	"go-project-frame/controller/user"
)

func InstallRouter(apiGroup *gin.RouterGroup) {
	installUnOperationRouter(apiGroup)
	installOperationRouter(apiGroup)
}

func installUnOperationRouter(apiGroup *gin.RouterGroup) {
	user.NewUserRouter(apiGroup)
}

func installOperationRouter(apiGroup *gin.RouterGroup) {

}
