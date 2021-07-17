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

func testCreate(t *testing.T) int64 {
	fmt.Println(`test /vips/create`)
	client := &http.Client{}
	postData := url.Values{}
	postData.Set(`name`, `test1`)
	postData.Set(`country`, `Indonesia`)
	postData.Set(`eta`, `2021-01-01 23:59:59`)
	// dst
	req, err := http.NewRequest("POST", "http://localhost"+config.API_PORTT+`/vips/create`, bytes.NewBufferString(postData.Encode()))
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(config.API_AUTH_USER, config.API_AUTH_PASS))
	req.Header.Add(`Content-Type`, `application/x-www-form-urlencoded`)
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return 0
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return 0
	}

	fmt.Println(string(bodyBytes))
	m := struct {
		VipId int64 `json:"data"`
		Error interface{}
	}{}
	err = json.Unmarshal(bodyBytes, &m)
	if err != nil {
		t.Error(err)
		return 0
	}
	return m.VipId
}

func TestVip(t *testing.T) {
	// start app
	go Run()

	// test create	
	id := testCreate(t)

	// test select
	if id == 0 {
		return
	}
	testGetById(t, id, false)

	// test arrive
	testArrive(t, id)

	// test check arrive
	testGetById(t, id, true)
}

func testArrive(t *testing.T, id int64) {
	idStr := strconv.Itoa(int(id))
	fmt.Println(`test /vips/arrive/` + idStr)
	client := &http.Client{}
	req, err := http.NewRequest("PATCH", "http://localhost"+config.API_PORTT+`/vips/arrive/`+idStr, nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(config.API_AUTH_USER, config.API_AUTH_PASS))
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(bodyBytes))

	res := struct {
		Data  interface{}
		Error string
	}{}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		t.Error(err)
		return
	}

	if res.Error != `` {
		t.Error(`arrive failed`)
		return
	}
}

func testGetById(t *testing.T, id int64, shouldArrive bool) {
	idStr := strconv.Itoa(int(id))
	fmt.Println(`test /vips/by-id/` + idStr)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost"+config.API_PORTT+`/vips/by-id/`+idStr, nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Add("Authorization", "Basic "+basicAuth(config.API_AUTH_USER, config.API_AUTH_PASS))
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(bodyBytes))
	res := struct {
		Data struct {
			ID         int64
			Name       string
			Country    string
			ETA        string
			Arrived    bool
			PhotoUrl   string
			Attributes string
		}
		Error interface{}
	}{}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res)

	if res.Data.ID != id {
		t.Error(`fetched VIP not exists`)
		return
	}

	if res.Data.Arrived != shouldArrive {
		t.Errorf(`fetched VIP arrived status should be %v`, shouldArrive)
		return
	}
}
