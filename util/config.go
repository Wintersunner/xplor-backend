package util

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
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
	viper.SetDefault("GIN_MODE", "release")
	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DB_DRIVER", "mysql")
	viper.SetDefault("DB_NAME", "default_db_name")
	viper.SetDefault("DB_USERNAME", "default_db_username")
	viper.SetDefault("DB_PASSWORD", "default_db_password")
	viper.SetDefault("DB_HOST", "default_db_host")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("ALLOWED_ORIGINS", "3306")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config file not found using os env variables")
		} else {
			log.Fatal("error loading config file", err)
			return
		}
	}

	err = viper.Unmarshal(&config)
	fmt.Println(config)
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
