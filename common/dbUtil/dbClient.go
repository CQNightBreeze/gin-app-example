package dbUtil

import (
	"fmt"
	"gin-app-example/common/util"
	"gin-app-example/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var DbClient *gorm.DB

func init() {
	envConfig := util.GetEnvConfig()
	strDBConnect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local", envConfig.Mysql.User,
		envConfig.Mysql.Password, envConfig.Mysql.Address, envConfig.Mysql.Port, envConfig.Mysql.TableName)

	logrus.Printf("strDBConnect:  %v", strDBConnect)
	db, err := gorm.Open("mysql", strDBConnect)
	if !envConfig.Env.Pro {
		db.LogMode(true)
	}
	if err != nil {
		//logrus.Printf("failed to connect database %v", err)
		logrus.Panicf("failed to connect database  %v", err)
	}

	defer db.Close()
	db.AutoMigrate(&models.User{})
	DbClient = db
}
