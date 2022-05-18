package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetLogPath(name string) string {
	if !strings.HasSuffix(name, ".log") {
		name = name + ".log"
	}
	envConfig := GetEnvConfig()
	var logDir string
	if envConfig.LogDir == "" {
		logDir = os.Getenv("LOG_DIR")
		if logDir == "" {
			cwd, _ := os.Getwd()
			logDir = filepath.Join(cwd, "runtime")
		}
	}
	outPath := filepath.Join(logDir, name)
	outPath = NormalizePath(outPath)
	return outPath
}
