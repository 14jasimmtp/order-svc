package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct{
	DB_URL string `mapstructure:"DB_URL"`
	PORT string  `mapstructure:"PORT"`
	PRODUCT_SVC_URL string `mapstructure:"PRODUCT_SVC_URL"`
}

func NewConfig () (config Config,err error){
	viper.AddConfigPath("./pkg/config/env")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil{
		log.Fatalln("err in viper reading ", err)
	}

	err =viper.Unmarshal(&config)

	return
}