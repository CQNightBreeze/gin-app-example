package apis

import (
	"github.com/gin-gonic/gin"

	user "gin-app-example/handlers/user"
)

func ClientApis(r *gin.Engine) {
	userHandler := user.UserHandler{}

	clientApi := r.Group("/c")

	clientApi.POST("/UserHandler/SignIn", userHandler.SignIn)
	clientApi.GET("/UserHandler/GetUserInfo", userHandler.GetUserInfo)
	clientApi.GET("/UserHandler/GetUserInfoConditional", userHandler.GetUserInfoConditional)

}
