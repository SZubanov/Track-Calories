package fatsecret

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	RequestTokenUrl = "https://www.fatsecret.com/oauth/request_token"
	AccessTokenUrl  = "https://www.fatsecret.com/oauth/access_token"
	ApiUrl          = "https://platform.fatsecret.com/rest/server.api"
)

type FatSecretConn struct {
	apikey string
	secret string
}

func Connect(apikey, secret string) (FatSecretConn, error) {
	return FatSecretConn{
		apikey,
		secret,
	}, nil
}

func (fs FatSecretConn) GetTokenMethods(oAuthSecret, url string, params map[string]string) (io.ReadCloser, error) {
	return fs.get(oAuthSecret, url, params)
}

func (fs FatSecretConn) GetApiMethods(oAuthSecret, method string, params map[string]string) (io.ReadCloser, error) {
	params["method"] = method
	params["format"] = "json"
	return fs.get(oAuthSecret, ApiUrl, params)
}

func (fs FatSecretConn) get(oAuthSecret, apiUrl string, params map[string]string) (io.ReadCloser, error) {
	reqTime := fmt.Sprintf("%d", time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := map[string]string{
		"oauth_consumer_key":     fs.apikey,
		"oauth_nonce":            fmt.Sprintf("%d", r.Int63()),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        reqTime,
		"oauth_version":          "1.0",
	}
	for k, v := range params {
		m[k] = v
	}
	mk := make([]string, len(m))
	i := 0
	for k, _ := range m {
		mk[i] = k
		i++
	}
	// sort keys
	sort.Strings(mk)

	// build sorted k/v string for sig
	sigQueryStr := ""
	for _, k := range mk {
		sigQueryStr += fmt.Sprintf("&%s=%s", k, escape(m[k]))
	}
	// drop initial &
	sigQueryStr = sigQueryStr[1:]
	sigBaseStr := fmt.Sprintf("GET&%s&%s", url.QueryEscape(apiUrl), escape(sigQueryStr))
	//fmt.Println("sigstr:", sigBaseStr)

	mac := hmac.New(sha1.New, []byte(fs.secret+"&"+oAuthSecret))
	mac.Write([]byte(sigBaseStr))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// add sig to map
	m["oauth_signature"] = sig
	mk = append(mk, "oauth_signature")

	// re-sort keys after adding sig
	sort.Strings(mk)
	requrl := fmt.Sprintf("%s?", apiUrl)
	reqQuery := ""
	for _, k := range mk {
		reqQuery += fmt.Sprintf("&%s=%s", k, escape(m[k]))
	}
	// drop initial &
	reqQuery = reqQuery[1:]

	requrl += reqQuery
	//fmt.Println("url :", requrl)
	resp, err := http.Get(requrl)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func escape(s string) string {
	return strings.Replace(strings.Replace(url.QueryEscape(s), "+", "%20", -1), "%7E", "~", -1)
}
