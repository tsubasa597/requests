package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	Client = &http.Client{}
	Header = map[string][]string{
		"Connection":   {"keep-alive"},
		"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
		"Content-Type": {"application/x-www-form-urlencoded"},
	}
	Cookie = make(map[string][]string)
)

func Get(url string) ([]byte, error) {
	var err error

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "地址错误", err)
	}

	req.Header = Header

	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "请求错误", err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "读取错误", err)
	}

	return res, err
}

func Gets(url string, v interface{}) error {
	resp, err := Get(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, v)
	if err != nil {
		return fmt.Errorf("%s: %s", "解析错误", err)
	}

	return nil
}

func Post(url string, params url.Values) ([]byte, error) {
	var err error

	req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "地址错误", err)
	}

	req.Header = Header

	resp, err := Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "请求错误", err)
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "读取错误", err)
	}

	return res, err
}

func Posts(url string, params url.Values, v interface{}) error {
	resp, err := Post(url, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, v)
	if err != nil {
		return fmt.Errorf("%s: %s", "解析错误", err)
	}

	return nil
}
