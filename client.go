package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct{}

func (c *Client) Get(url string, params url.Values, r interface{}) error {
	req, err := http.NewRequest("GET", url+"?"+params.Encode(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Post(url string, payload url.Values, r interface{}) error {
	data, err := c.PostRaw(url, payload)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) PostRaw(url string, payload url.Values) ([]byte, error) {
	body := strings.NewReader(payload.Encode())
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	return data, err
}
