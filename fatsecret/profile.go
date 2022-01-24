package fatsecret

import (
	"fmt"
	"io/ioutil"
)

func (fs FatSecretConn) GetAuth(query string) (interface{}, error) {
	resp, err := fs.get(
		"profile.get_auth",
		map[string]string{"user_id": query},
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
