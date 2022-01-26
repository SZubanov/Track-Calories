package fatsecret

import (
	"encoding/json"
	"errors"
	"github.com/SZubanov/Track-Calories/helpers"
	"io/ioutil"
)

type FoodEntriesResponse struct {
	FoodEntries FoodEntries    `json:"food_entries"`
	Error       *ErrorResponse `json:"error"`
}

type FoodEntries struct {
	FoodEntry []FoodEntry `json:"food_entry"`
}

type FoodEntry struct {
	ID           string `json:"food_id"`
	Date         string `json:"date_int"`
	Calories     string `json:"calories"`
	Carbohydrate string `json:"carbohydrate"`
	Protein      string `json:"protein"`
	Fat          string `json:"fat"`
	Fiber        string `json:"fiber"`
}

func (fs FatSecretConnect) GetFoodEntry() (*FoodEntriesResponse, error) {
	resp, err := fs.GetApiMethods(
		"food_entries.get",
		map[string]interface{}{
			"date": helpers.GetYesterdayUnix(),
		},
	)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	foodresp := FoodEntriesResponse{}
	err = json.Unmarshal(body, &foodresp)
	if err != nil {
		return nil, err
	}
	if foodresp.Error != nil {
		return nil, errors.New(foodresp.Error.Message)
	}
	return &foodresp, nil
}
