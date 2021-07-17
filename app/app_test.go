package app

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"vipms1/app/config"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func testInsert(t *testing.T) {
	client := &http.Client{}
	postData := url.Values{}
	postData.Set(`name`, `test1`)
	postData.Set(`country`, `Indonesia`)
	postData.Set(`eta`, `2021-01-01 23:59:59`)
	// dst
	req, err := http.NewRequest("POST", "http://localhost"+config.API_PORTT+`/vips/create`, bytes.NewBuffer(postData))
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(config.API_AUTH_USER, config.API_AUTH_PASS))
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(bodyBytes))
	m := map[string]interface{}
	err := json.Unmarshal(bodyBytes, &m)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(`data = `, m[`data`])
	m2, ok := m[`data`].(map[string]interface{})
	if ok {
		fmt.Println(`id = `, m2[`id`])
	}
}

func TestVip(t *testing.T) {
	// test insert	

	// test select
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost"+config.API_PORTT+`/vips/by-id/`+strconv.Itoa(id), nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(config.API_AUTH_USER, config.API_AUTH_PASS))
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(bodyBytes))
}
