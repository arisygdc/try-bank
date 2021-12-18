package config

import "github.com/spf13/viper"

type Environment struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESSESS"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
}

func NewEnv(path string) (env Environment, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&env)
	return
}
