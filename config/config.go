package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("fail to get current file path")
	}

	dir := filepath.Dir(filename)

	configPath := filepath.Join(dir, "config.yaml")

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic("fail to read config file:" + err.Error())
	}
}
