package configs

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *appConfig

// Initilize the configuration with environment variables.
func Init() {
	AppConfig = loadConfig()
}

type appConfig struct {
	App               string `mapstructure:"APP"`
	LineAccessToken   string `mapstructure:"LINE_ACCESS_TOKEN"`
	LineChannelSecret string `mapstructure:"LINE_CHANNEL_SECRET"`
	MongoDB           string `mapstructure:"MONGO_DB"`
	MongoUri          string `mapstructure:"MONGO_URI"`
	NgrokAuthToken    string `mapstructure:"NGROK_AUTHTOKEN"`
}

func loadConfig() *appConfig {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading environment file", err)
	}
	var config *appConfig
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return config
}
