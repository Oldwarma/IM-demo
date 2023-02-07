package main

import (
	"github.com/spf13/viper"
)

func main() {
	NewSetting()
	for {

	}
}

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigName("im")
	vp.AddConfigPath("etc/")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
