package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBPort        string `mapstructure:"DB_PORT"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// reads the configuration from the environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
