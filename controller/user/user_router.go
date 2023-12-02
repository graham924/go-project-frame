package user

import "github.com/gin-gonic/gin"

type userController struct {
}

func NewUserRouter(ginEngine *gin.RouterGroup) {
	u := userController{}
	u.initRoutes(ginEngine)
}

func (u *userController) initRoutes(ginEngine *gin.RouterGroup) {
	userRoute := ginEngine.Group("/user")
	user := userController{}
	{
		userRoute.POST("/login", user.Login)
	}
}
