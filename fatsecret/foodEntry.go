package fatsecret

import (
	"fmt"
	"io/ioutil"
)

func (fs FatSecretConn) GetFoodEntryMonth() (interface{}, error) {
	resp, err := fs.GetApiMethods(
		"a432f0eefc2543ca8f5877cc17e1f107",
		"food_entries.get_month",
		map[string]string{
			"oauth_token": "11faf9679618496cb5d3e05ecfd7e206",
		},
	)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	fmt.Println(string(body))
	//foodresp := FoodSearchResponse{}
	//err = json.Unmarshal(body, &foodresp)
	//if err != nil {
	//	return nil, err
	//}
	//if foodresp.Error != nil {
	//	return nil, errors.New(foodresp.Error.Message)
	//}
	return body, nil
}

func (fs FatSecretConn) GetFoodEntry() (interface{}, error) {
	resp, err := fs.GetApiMethods(
		"a432f0eefc2543ca8f5877cc17e1f107",
		"food_entries.get",
		map[string]string{
			"oauth_token": "11faf9679618496cb5d3e05ecfd7e206",
		},
	)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	fmt.Println(string(body))
	//foodresp := FoodSearchResponse{}
	//err = json.Unmarshal(body, &foodresp)
	//if err != nil {
	//	return nil, err
	//}
	//if foodresp.Error != nil {
	//	return nil, errors.New(foodresp.Error.Message)
	//}
	return body, nil
}
