package main

import (
	"log"

	"github.com/cyansobble/global"
	"github.com/spf13/viper"
)

func Viper() {
	viper.SetConfigFile("./config.yaml")
	// viper.SetConfigName("config")
	// viper.SetConfigType("ymal")
	// viper.AddConfigPath("")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//log.Errorf()
			log.Fatalln("viper config file not found", err)
		} else {
			log.Fatalln("viper", err)
		}
	} else {
		log.Println("viper read config success")
	}

	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		log.Fatalln("unmarshal conf failed", err)
	}

}
