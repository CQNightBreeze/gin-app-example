package util

// util 内部工具包
import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

// EnvConfig 根据运行环境需要实际配置的值，包括数据库地址，表名这些
type EnvConfig struct {
	Mysql    Mysql   `json:"mysql"`
	Redis    Redis   `json:"redis"`
	Mongodb  Mongodb `json:"mongodb"`
	Env      Env     `json:"env"`
	LogIsOut bool    `json:"logIsOut"`
	LogDir   string  `json:"logDir"`
}

type Env struct {
	Runtime string `json:"runtime"` // “dev”是开发环境，"prod"代表生产环节
	Pro     bool   `json:"pro"`     // false 是开发环境，true 代表生产环节代表生产环节
}

type Mongodb struct {
	Address  string `json:"address"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	TableName string `json:"tableName"`
}
type Redis struct {
	Address  string `json:"address"`
	Port     uint   `json:"port"`
	Password string `json:"password"`
}

type Mysql struct {
	Address   string `json:"address"`
	Port      uint   `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	TableName string `json:"tableName"`
}

var once sync.Once

// getEnvConfigInstance 所有的配置，都是从这个单例封装，后期可实现动态配置
func getEnvConfigInstance() *EnvConfig {
	var v EnvConfig
	once.Do(func() {
		data, err := ioutil.ReadFile(EnvConfigFile)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(data, &v)
		if err != nil {
			panic(err)
		}
	})
	return &v
}

// GetEnvConfig 所有的配置，都是从这个单例翻转，后期可实现动态配置
func GetEnvConfig() *EnvConfig {
	return getEnvConfigInstance()
}
