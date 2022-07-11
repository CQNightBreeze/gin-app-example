package apis

import (
	admin "gin-app-example/manage_handlers/admin"

	"github.com/gin-gonic/gin"
)

func ManageApis(r *gin.Engine) {
	manageApi := r.Group("/m")

	adminHandler := admin.AdminHandler{}

	manageApi.GET("/AdminHandler/SignIn", adminHandler.SignIn)

}
