package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppName        string `mapstructure:"APP_ENV"`
	AppURL         string `mapstructure:"APP_URL"`
	GinMode        string `mapstructure:"GIN_MODE"`
	Host           string `mapstructure:"HOST"`
	Port           string `mapstructure:"PORT"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBName         string `mapstructure:"DB_NAME"`
	DBUsername     string `mapstructure:"DB_USERNAME"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	AllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.SetDefault("ALLOWED_ORIGINS", "*")
	viper.SetDefault("GIN_MODE", "release")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func (config *Config) DBSource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DBUsername, config.DBPassword,
		config.DBHost, config.DBPort, config.DBName)
}

func (config *Config) MigrationSource() string {
	return fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s", config.DBDriver, config.DBUsername, config.DBPassword,
		config.DBHost, config.DBPort, config.DBName)
}
