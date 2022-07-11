package util

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// EnvConfigFile 环境配置
	EnvConfigFile = GetConfigFile(".envconfig.json")
	// RbacModelFile 权限配置
	RbacModelFile = GetConfigFile("accessctl/rbac_model.conf")
	// RbacPolicyFile 权限配置
	RbacPolicyFile = GetConfigFile("rbac_policy.csv")
)

// NormalizePath 路径转换
func NormalizePath(path string) string {
	switch runtime.GOOS {
	case "darwin":
	case "windows":
		// path = strings.Replace(path, "\\", "\\\\", -1)
		path = strings.Replace(path, "\\", "/", -1)
		break
	case "linux":
	}
	return path
}

// GetConfigFile 获取坏境文件
func GetConfigFile(file string) string {
	configDir := os.Getenv("CONFIG_DIR")
	if configDir == "" {
		cwd, _ := os.Getwd()
		configDir = filepath.Join(cwd, "config")
	}
	return NormalizePath(filepath.Join(configDir, file))
}
