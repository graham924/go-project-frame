package user

import (
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOut} "success"
// @Router /api/user/login [post]
func (u *userController) Login(ctx *gin.Context) {
	ctx.JSON(200, "success")
}
