package main

import (
	"log"

	"github.com/datshiro/cyclo-ecommerce/internal/infras/server"
	"github.com/datshiro/cyclo-ecommerce/internal/infras/server/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	pflag.String("env", "staging", "Env mode")
	pflag.String("configPath", "./", "Env directory path")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	config, err := config.NewConfig(
		viper.GetString("configPath"),
		viper.GetString("env"),
	)
	if err != nil {
		log.Fatalf("Failed to init configuration: %v", err)
	}
	log.Printf("Config %v", config)

	srv := server.New(config)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
