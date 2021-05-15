package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Requests struct {
	header map[string][]string
	req    *http.Request
	cli    *http.Client
}

func (r *Requests) Get(url string) ([]byte, error) {
	var err error

	if len(r.header) < 1 {
		r.header = map[string][]string{
			"Connection":   {"keep-alive"},
			"User-Agent":   {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
			"Content-Type": {"application/x-www-form-urlencoded"},
		}
	}

	r.req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "地址错误", err)
	}

	r.req.Header = r.header

	resp, err := r.cli.Do(r.req)
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

func (r *Requests) Gets(url string, v *interface{}) error {
	resp, err := r.Get(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, v)
	if err != nil {
		return fmt.Errorf("%s: %s", "解析错误", err)
	}

	return nil
}

func (r *Requests) Post(url string, params url.Values) ([]byte, error) {
	var err error

	if len(r.header) < 1 {
		r.req.Header = http.Header{
			"Connection":   []string{"keep-alive"},
			"User-Agent":   []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36 Edg/85.0.564.70"},
			"Content-Type": []string{"application/x-www-form-urlencoded"},
		}
	}

	r.req, err = http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("%s: %s", "地址错误", err)
	}

	r.req.Header = r.header

	resp, err := r.cli.Do(r.req)
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

func (r *Requests) Posts(url string, params url.Values, v *interface{}) error {
	resp, err := r.Post(url, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, v)
	if err != nil {
		return fmt.Errorf("%s: %s", "解析错误", err)
	}

	return nil
}

func (r *Requests) SetHeader(h http.Header) {
	r.header = h
}

func New() *Requests {
	return &Requests{
		req:    &http.Request{},
		cli:    &http.Client{},
		header: http.Header{},
	}
}
