package httpclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"transformer_mz/libs/errors"
)

type HTTPClient struct {
	token       string
	baseURL     string
	contentType string
}

func NewClient(args ...string) *HTTPClient {
	client := &HTTPClient{
		baseURL:     "http://localhost:8080",
		contentType: "application/json; charset=UTF-8",
	}
	if len(args) > 0 {
		client.token = args[0]
	}
	if len(args) > 1 {
		client.baseURL = args[1]
	}
	if len(args) > 2 {
		client.contentType = args[2]
	}
	return client
}

func (c *HTTPClient) process(url, method string, beanIn interface{}) (io.ReadCloser, error) {
	var (
		data []byte
		err  error
	)
	switch beanIn.(type) {
	case []byte:
		data = beanIn.([]byte)
	case string:
		data = []byte(beanIn.(string))
	default:
		data, err = json.Marshal(beanIn)
		if err != nil {
			return nil, errors.New(err, "服务间调用失败，请联系管理员进行处理")
		}
	}
	req, err := http.NewRequest(method, c.baseURL+url, bytes.NewReader(data))
	if err != nil {
		return nil, errors.New(err, "服务间调用失败，请联系管理员进行处理")
	}
	req.Header.Set("Content-Type", c.contentType)
	req.Header.Set("Token", c.token)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err, "服务间调用失败，请联系管理员进行处理")
	}
	if resp.StatusCode != http.StatusOK {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(err, "服务间调用失败，请联系管理员进行处理")
		}
		return nil, errors.New(string(data), "服务间调用失败，请联系管理员进行处理")
	}
	return resp.Body, nil
}

func (c *HTTPClient) GetSpec(url string, beanIn interface{}) (io.ReadCloser, error) {
	return c.process(url, http.MethodGet, beanIn)
}

func (c *HTTPClient) PutSpec(url string, beanIn interface{}) (io.ReadCloser, error) {
	return c.process(url, http.MethodPut, beanIn)
}

func (c *HTTPClient) PostSpec(url string, beanIn interface{}) (io.ReadCloser, error) {
	return c.process(url, http.MethodPost, beanIn)
}

func (c *HTTPClient) PatchSpec(url string, beanIn interface{}) (io.ReadCloser, error) {
	return c.process(url, http.MethodPatch, beanIn)
}

func (c *HTTPClient) DeleteSpec(url string, beanIn interface{}) (io.ReadCloser, error) {
	return c.process(url, http.MethodDelete, beanIn)
}
