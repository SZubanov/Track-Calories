package config

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	FatSecretApiKey   string `json:"fatsecret_apikey" mapstructure:"fatsecret_apikey" validate:"required"`
	FatSecretSecret   string `json:"fatsecret_secret" mapstructure:"fatsecret_secret" validate:"required"`
	OAuthToken        string `json:"oauth_token" mapstructure:"oauth_token" validate:"required"`
	OAuthTokenSecret  string `json:"oauth_token_secret" mapstructure:"oauth_token_secret" validate:"required"`
	FormUrl           string `json:"form_url" mapstructure:"form_url" validate:"required"`
	DayInput          string `json:"dayInput" mapstructure:"day_input" validate:"required"`
	MonthInput        string `json:"monthInput" mapstructure:"month_input" validate:"required"`
	YearInput         string `json:"yearInput" mapstructure:"year_input" validate:"required"`
	WeightInput       string `json:"weight" mapstructure:"weight_input" validate:"required"`
	CaloriesInput     string `json:"calories" mapstructure:"calories_input" validate:"required"`
	ProteinInput      string `json:"protein" mapstructure:"protein_input" validate:"required"`
	FatInput          string `json:"fat" mapstructure:"fat_input" validate:"required"`
	CarbohydrateInput string `json:"carbohydrate" mapstructure:"carbohydrate_input" validate:"required"`
	FiberInput        string `json:"fiber" mapstructure:"fiber_input" validate:"required"`
	WaterInput        string `json:"water" mapstructure:"water_input" validate:"required"`
}

// NewConfig load current configuration.
func NewConfig() (*Config, error) {
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("NewConfig(): failed to unmarshaled %w", err)
	}
	err := validator.New().Struct(config)
	if err != nil {
		return nil, fmt.Errorf("NewConfig(): failed to validated %w", err)
	}

	return &config, nil
}
