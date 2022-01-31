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
	FatSecretApiKey   string `json:"fatsecret_apikey" mapstructure:"fatsecret_apikey"`
	FatSecretSecret   string `json:"fatsecret_secret" mapstructure:"fatsecret_secret"`
	OAuthToken        string `json:"oauth_token" mapstructure:"oauth_token"`
	OAuthTokenSecret  string `json:"oauth_token_secret" mapstructure:"oauth_token_secret"`
	FormUrl           string `json:"form_url" mapstructure:"form_url"`
	DayInput          string `json:"dayInput" mapstructure:"day_input"`
	MonthInput        string `json:"monthInput" mapstructure:"month_input"`
	YearInput         string `json:"yearInput" mapstructure:"year_input"`
	WeightInput       string `json:"weight" mapstructure:"weight_input"`
	CaloriesInput     string `json:"calories" mapstructure:"calories_input"`
	ProteinInput      string `json:"protein" mapstructure:"protein_input"`
	FatInput          string `json:"fat" mapstructure:"fat_input"`
	CarbohydrateInput string `json:"carbohydrate" mapstructure:"carbohydrate_input"`
	FiberInput        string `json:"fiber" mapstructure:"fiber_input"`
	WaterInput        string `json:"water" mapstructure:"water_input"`
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
