package main

import (
	"gin-app-example/common/logs"
	docs "gin-app-example/docs"

	apis "gin-app-example/apis"

	"gin-app-example/common/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	envConfig := util.GetEnvConfig()
	logrus.SetFormatter(&logrus.TextFormatter{ /*  */
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	if envConfig.LogIsOut {
		hook := logs.NewHook(util.GetLogPath("logs"))
		logrus.AddHook(hook)
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Host = ""

	apis.ClientApis(r)
	apis.ManageApis(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
