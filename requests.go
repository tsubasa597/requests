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
	defaultHeaders = map[string]string{
		"Connection":   "keep-alive",
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70",
		"Content-Type": "application/x-www-form-urlencoded",
	}
	request = &Requests{
		Client: &http.Client{},
	}
)

type Requests struct {
	Client  *http.Client
	Headers map[string]string
	Cookies map[string]string
}

func (request Requests) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "地址错误", err)
	}

	setHeadersAndCookies(request.Headers, request.Cookies, req)

	resp, err := request.Client.Do(req)
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

func (request Requests) Gets(url string, v interface{}) error {
	resp, err := request.Get(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, v)
	if err != nil {
		return fmt.Errorf("%s: %s", "解析错误", err)
	}

	return nil
}

func (request Requests) Post(url string, params url.Values) ([]byte, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "地址错误", err)
	}

	setHeadersAndCookies(request.Headers, request.Cookies, req)

	resp, err := request.Client.Do(req)
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

func (request Requests) Posts(url string, params url.Values, v interface{}) error {
	resp, err := request.Post(url, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, v)
	if err != nil {
		return fmt.Errorf("%s: %s", "解析错误", err)
	}

	return nil
}

func setHeadersAndCookies(headers, cookies map[string]string, req *http.Request) {
	if headers == nil || len(headers) < 1 {
		headers = defaultHeaders
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	for k, v := range cookies {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

func Get(url string) ([]byte, error) {
	return request.Get(url)
}

func Post(url string, params url.Values) ([]byte, error) {
	return request.Post(url, params)
}
