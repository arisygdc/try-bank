package config

import "github.com/spf13/viper"

type Environment struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESSESS"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	Environment   string `mapstructure:"ENVIRONMENT"`
}

// Supply param for configuration from .env file or environment variable if exists
// default config file name config.env

// configName param when to use custom name
// example: file name is app.env
// fill "app" in configName param
func NewEnv(path string, configName ...string) (env Environment, err error) {
	viper.AddConfigPath(path)
	if len(configName) == 1 {
		viper.SetConfigName(configName[0])
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&env)
	return
}
