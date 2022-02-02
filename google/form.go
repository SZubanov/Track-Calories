package google

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ResultRequest struct {
	Day          int     `json:"day"`
	Month        int     `json:"month"`
	Year         int     `json:"year"`
	Weight       float64 `json:"weight"`
	Calories     float64 `json:"calories"`
	Carbohydrate float64 `json:"carbohydrate"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Fiber        float64 `json:"fiber"`
	Water        float64 `json:"water"`
}

func NewResultRequest(day, month, year int, weight, calories, carbohydrate, protein, fat, fiber, water float64) *ResultRequest {
	return &ResultRequest{
		Day:          day,
		Month:        month,
		Year:         year,
		Weight:       weight,
		Calories:     calories,
		Carbohydrate: carbohydrate,
		Protein:      protein,
		Fat:          fat,
		Fiber:        fiber,
		Water:        water,
	}
}

func (r ResultRequest) ToRequestMap(fields FormFields) map[string]string {
	return map[string]string{
		fields.Day:          strconv.Itoa(r.Day),
		fields.Month:        strconv.Itoa(r.Month),
		fields.Year:         strconv.Itoa(r.Year),
		fields.Weight:       fmt.Sprintf("%f", r.Weight),
		fields.Calories:     fmt.Sprintf("%f", r.Calories),
		fields.Protein:      fmt.Sprintf("%f", r.Protein),
		fields.Carbohydrate: fmt.Sprintf("%f", r.Carbohydrate),
		fields.Fat:          fmt.Sprintf("%f", r.Fat),
		fields.Fiber:        fmt.Sprintf("%f", r.Fiber),
		fields.Water:        fmt.Sprintf("%f", r.Water),
	}
}

type FormFields struct {
	Day          string `json:"day"`
	Month        string `json:"month"`
	Year         string `json:"year"`
	Weight       string `json:"weight"`
	Calories     string `json:"calories"`
	Carbohydrate string `json:"carbohydrate"`
	Protein      string `json:"protein"`
	Fat          string `json:"fat"`
	Fiber        string `json:"fiber"`
	Water        string `json:"water"`
}

func NewFormFields(day, month, year, weight, calories, carbohydrate, protein, fat, fiber, water string) *FormFields {
	return &FormFields{
		Day:          day,
		Month:        month,
		Year:         year,
		Weight:       weight,
		Calories:     calories,
		Carbohydrate: carbohydrate,
		Protein:      protein,
		Fat:          fat,
		Fiber:        fiber,
		Water:        water,
	}
}

func RequestForm(requestUrl string, fields FormFields, params ResultRequest) (*http.Response, error) {
	body := params.ToRequestMap(fields)
	values := url.Values{}
	for key, value := range body {
		values.Set(key, value)
	}

	resp, err := http.PostForm(requestUrl, values)
	return resp, err
}
