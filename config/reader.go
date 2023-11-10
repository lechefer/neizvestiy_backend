package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Read() Config {
	env, configName := readEnv()

	viper.SetConfigName(configName) // name of config file (without extension)
	viper.SetConfigType("json")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // path to look for the config file in

	// Find and read the config file
	// Handle error reading the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var c Config
	c.Env = env
	if err := viper.Unmarshal(&c); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return c
}

const _serviceEnvVarName = "ENV"

type Env string

const (
	EnvProd  = Env("prod")
	EnvTest  = Env("test")
	EnvLocal = Env("local")
	EnvEmpty = Env("")
)

func readEnv() (Env, string) {
	env := Env(os.Getenv(_serviceEnvVarName))

	var configName string
	switch env {
	case EnvProd:
		configName = "prod"
	case EnvTest:
		configName = "test"
	default:
		env = EnvLocal
		configName = "local"
	}

	return env, configName
}
