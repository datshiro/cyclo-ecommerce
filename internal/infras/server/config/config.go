package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func NewConfig(path string, env string) (config Config, err error) {
	envFile := fmt.Sprintf("%s.env", strings.ToLower(env))
	fmt.Println("Loading config from ", filepath.Join(path, envFile))

	viper.AddConfigPath(path)
	viper.SetConfigName(
		envFile,
	)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

type Config struct {
	Debug         bool
	Host          string
	Port          string
	JwtSecretKey  string
	LogLevel      string
	DbUrl         string
	RedisLocation string
	ApiPath       string
}

func (c Config) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func Default() Config {
	cfg := Config{
		Debug:         true,
		Host:          "localhost",
		Port:          "8080",
		JwtSecretKey:  "jwt_siecret_key",
		LogLevel:      "DEBUG",
		DbUrl:         "postgres://postgres:postgres@localhost:54321/etl?sslmode=disable",
		RedisLocation: "",
		ApiPath:       "/api",
	}
	d1 := []byte("")
	v := reflect.ValueOf(cfg)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			d1 = append(
				d1,
				fmt.Sprintf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())...,
			)
		case reflect.Struct:
			iterateConfigFields(v.Field(i).Interface(), d1)
		}
	}

	err := os.WriteFile("example.env.yaml", d1, 0644)
	if err != nil {
		log.Printf("Failed to write example env file: %v", err)
	}
	return cfg
}

func iterateConfigFields(cfg interface{}, data []byte) {
	v := reflect.ValueOf(cfg)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			data = append(
				data,
				fmt.Sprintf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())...,
			)
		case reflect.Struct:
			iterateConfigFields(v.Field(i).Interface(), data)
		}
	}
}
