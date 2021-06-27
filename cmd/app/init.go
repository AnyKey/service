package main

import (
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("cannot read config", err)
	}
	return
}

func initLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	if viper.GetBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	return
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.port"),
	})
	return rdb
}


