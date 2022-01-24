package fatsecret

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"time"
)

var requestUrl = "https://www.fatsecret.com/oauth/request_token"

func (fs FatSecretConn) getRequestToken() (io.ReadCloser, error) {
	reqTime := fmt.Sprintf("%d", time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := map[string]string{
		"oauth_consumer_key":     fs.apikey,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        reqTime,
		"oauth_nonce":            fmt.Sprintf("%d", r.Int63()),
		"oauth_version":          "1.0",
		"oauth_callback":         "oob",
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
	sigBaseStr := fmt.Sprintf("GET&%s&%s", url.QueryEscape(requestUrl), escape(sigQueryStr))
	//fmt.Println("sigstr:", sigBaseStr)

	mac := hmac.New(sha1.New, []byte(fs.secret+"&"))
	mac.Write([]byte(sigBaseStr))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// add sig to map
	m["oauth_signature"] = sig
	mk = append(mk, "oauth_signature")

	// re-sort keys after adding sig
	sort.Strings(mk)
	requrl := fmt.Sprintf("%s?", requestUrl)
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

func (fs FatSecretConn) RequestToken() (interface{}, error) {
	resp, err := fs.getRequestToken()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	fmt.Println(string(body))
	return body, nil
}

var accessUrl = "https://www.fatsecret.com/oauth/access_token"

func (fs FatSecretConn) getAuthToken() (io.ReadCloser, error) {
	reqTime := fmt.Sprintf("%d", time.Now().Unix())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := map[string]string{
		"oauth_consumer_key":     fs.apikey,
		"oauth_token":            "14273763c8a943c0b9f97a9cfe01c1f7",
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        reqTime,
		"oauth_nonce":            fmt.Sprintf("%d", r.Int63()),
		"oauth_version":          "1.0",
		"oauth_verifier":         "9351096",
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
	sigBaseStr := fmt.Sprintf("GET&%s&%s", url.QueryEscape(accessUrl), escape(sigQueryStr))
	//fmt.Println("sigstr:", sigBaseStr)

	mac := hmac.New(sha1.New, []byte(fs.secret+"&88b8c303e4d049cf9d2294ebaaa9dd18"))
	mac.Write([]byte(sigBaseStr))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// add sig to map
	m["oauth_signature"] = sig
	mk = append(mk, "oauth_signature")

	// re-sort keys after adding sig
	sort.Strings(mk)
	requrl := fmt.Sprintf("%s?", accessUrl)
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

func (fs FatSecretConn) AuthToken() (interface{}, error) {
	resp, err := fs.getAuthToken()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp)
	defer resp.Close()
	fmt.Println(string(body))
	return body, nil
}
