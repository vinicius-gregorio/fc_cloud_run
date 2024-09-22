package config

import "github.com/spf13/viper"

type EnvConfig struct {
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
	WeatherAPIURL string `mapstructure:"WHEATHER_API_URL"`
	CEPAPIURL     string `mapstructure:"CEP_API_URL"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

func LoadConfig(path string) (*EnvConfig, error) {
	var cfg *EnvConfig

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
