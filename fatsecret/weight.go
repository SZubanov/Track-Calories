package fatsecret

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type MonthWeightResponse struct {
	Month Month          `json:"month"`
	Error *ErrorResponse `json:"error"`
}

type Month struct {
	Day []Day `json:"day"`
}

type Day struct {
	Date   string `json:"date_int"`
	Weight string `json:"weight_kg"`
}

func (fs FatSecretConnect) GetMonthWeight() (*MonthWeightResponse, error) {
	resp, err := fs.GetApiMethods(
		"weights.get_month",
		map[string]interface{}{},
	)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	weight := MonthWeightResponse{}
	err = json.Unmarshal(body, &weight)
	if err != nil {
		return nil, err
	}
	if weight.Error != nil {
		return nil, errors.New(weight.Error.Message)
	}
	return &weight, nil
}
