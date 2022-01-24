package fatsecret

import (
	"fmt"
	"io/ioutil"
)

func (fs FatSecretConn) GetRequestToken() (interface{}, error) {
	resp, err := fs.GetTokenMethods("", RequestTokenUrl, map[string]string{
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

func (fs FatSecretConn) GetAccessToken(oAuthToken, oAuthSecret, verifierCode string) (interface{}, error) {
	resp, err := fs.GetTokenMethods(oAuthSecret, AccessTokenUrl, map[string]string{
		"oauth_token":    oAuthToken,
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
