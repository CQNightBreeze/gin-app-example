package mgoUtil

import (
	"fmt"
	"gin-app-example/common/util"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

var MgDB *mgo.Database

func init() {
	envConfig := util.GetEnvConfig()

	strMgoDBConnect := fmt.Sprintf("%s:%d", envConfig.Mongodb.Address, envConfig.Mongodb.Port)
	logrus.Printf("mongo address  %v", strMgoDBConnect)
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{strMgoDBConnect},
		Direct:    false,
		Timeout:   time.Second * 1,
		Username:  envConfig.Mongodb.User,
		Password:  envConfig.Mongodb.Password,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	mgSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logrus.Panicf("connect mgo error %v", err)
		panic("connect mongo error")
	}
	mgSession.SetSocketTimeout(1 * time.Hour)
	defer mgSession.Close()
	mgDB := mgSession.DB(envConfig.Mongodb.TableName)

	MgDB = mgDB
}
