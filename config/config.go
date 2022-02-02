package config

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

func init() {
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
	DayInput          string `json:"day_input" mapstructure:"day_input" validate:"required"`
	MonthInput        string `json:"month_input" mapstructure:"month_input" validate:"required"`
	YearInput         string `json:"year_input" mapstructure:"year_input" validate:"required"`
	WeightInput       string `json:"weight_input" mapstructure:"weight_input" validate:"required"`
	CaloriesInput     string `json:"calories_input" mapstructure:"calories_input" validate:"required"`
	ProteinInput      string `json:"protein_input" mapstructure:"protein_input" validate:"required"`
	FatInput          string `json:"fat_input" mapstructure:"fat_input" validate:"required"`
	CarbohydrateInput string `json:"carbohydrate_input" mapstructure:"carbohydrate_input" validate:"required"`
	FiberInput        string `json:"fiber_input" mapstructure:"fiber_input" validate:"required"`
	WaterInput        string `json:"water_input" mapstructure:"water_input" validate:"required"`
}

// NewConfig load current configuration.
func NewConfig() (*Config, error) {
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("NewConfig(): failed to unmarshaled %w", err)
	}

	v := reflect.ValueOf(&config).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		fmt.Println(os.LookupEnv(strings.ToUpper(tag)))

		if val, ok := os.LookupEnv(strings.ToUpper(tag)); ok {
			field := v.Field(i)
			if field.CanAddr() && field.CanSet() {
				v.Field(i).SetString(val)
			}
		}
	}

	err := validator.New().Struct(config)
	if err != nil {
		return nil, fmt.Errorf("NewConfig(): failed to validated %w", err)
	}

	return &config, nil
}
