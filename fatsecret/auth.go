package fatsecret

import (
	"fmt"
	"io/ioutil"
)

func (fs FatSecretConnect) GetRequestToken() (interface{}, error) {
	resp, err := fs.GetTokenMethods(RequestTokenUrl, map[string]interface{}{
		"oauth_callback": "oob",
	})

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	fmt.Println(string(body))
	return body, nil
}

func (fs FatSecretConnect) GetAccessToken(verifierCode string) (interface{}, error) {
	resp, err := fs.GetTokenMethods(AccessTokenUrl, map[string]interface{}{
		"oauth_token":    fs.oauthToken,
		"oauth_verifier": verifierCode,
	})

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	fmt.Println(string(body))
	return body, nil
}
