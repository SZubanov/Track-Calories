package config

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

func init() {
	{
		c := Config{}

		bs, err := json.Marshal(&c)
		if err != nil {
			panic("unable to marshal config to json")
		}

		viper.AutomaticEnv()
		viper.SetConfigType("json")
		if err := viper.ReadConfig(bytes.NewBuffer(bs)); err != nil {
			panic("failed to set viper config")
		}
	}

	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	_ = viper.ReadInConfig()
}

type Config struct {
	FatSecretApiKey  string `json:"fatsecret_apikey" mapstructure:"fatsecret_apikey"`
	FatSecretSecret  string `json:"fatsecret_secret" mapstructure:"fatsecret_secret"`
	OAuthToken       string `json:"oauth_token" mapstructure:"oauth_token"`
	OAuthTokenSecret string `json:"oauth_token_secret" mapstructure:"oauth_token_secret"`
}

// NewConfig load current configuration.
func NewConfig() (*Config, error) {
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	err := validator.New().Struct(config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
