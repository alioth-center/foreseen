package entity

import (
	"github.com/alioth-center/infrastructure/config"
)

var (
	GlobalConfig Config
)

func init() {
	err := config.LoadConfig(&GlobalConfig, "config/config.yaml")
	if err != nil {
		panic(err)
	}
}

type Config struct {
	AppID     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
	Token     string `yaml:"token"`
	LogDir    string `yaml:"log_dir"`
}
