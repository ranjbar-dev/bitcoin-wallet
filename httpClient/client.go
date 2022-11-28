package httpClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func (httpClient *httpClient) SimpleGet(url string, res interface{}) error {

	requestHeaders := map[string]string{
		"User-Agent": UserAgent,
		"Accept":     "Application/json",
	}

	httpResp, _, status, err := httpClient.HttpGet(url, requestHeaders)

	if err != nil {
		return err
	}

	if status != 200 {
		return errors.New("status code :" + strconv.Itoa(status) + "-" + string(httpResp))
	}

	err = json.Unmarshal(httpResp, res)

	if err != nil {
		return err
	}

	return nil
}

func (httpClient *httpClient) SimplePost(url string, body interface{}, res interface{}) error {

	requestHeaders := map[string]string{
		"User-Agent": UserAgent,
		"Accept":     "Application/json",
	}

	httpResp, _, status, err := httpClient.HttpPost(url, body, requestHeaders)

	if err != nil {
		return err
	}

	if status != 200 {
		return errors.New("status code :" + strconv.Itoa(status) + "-" + string(httpResp))
	}

	if res != nil {
		err = json.Unmarshal(httpResp, res)

		if err != nil {
			return err
		}
	}

	return nil
}

func (httpClient *httpClient) HttpGet(url string, headers map[string]string) ([]byte, http.Header, int, error) {
	respBody := []byte("")
	header := http.Header{}
	statusCode := http.StatusBadRequest

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return respBody, header, statusCode, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.Client{
		Timeout: 40 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return respBody, header, statusCode, err
	}
	defer resp.Body.Close()
	header = resp.Header
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return respBody, header, statusCode, err
	}
	statusCode = resp.StatusCode

	return respBody, header, statusCode, nil

}

func (httpClient *httpClient) HttpPost(url string, body interface{}, headers map[string]string) ([]byte, http.Header, int, error) {
	finalRes := []byte("")
	statusCode := http.StatusBadRequest
	header := http.Header{}
	requestBody, err := json.Marshal(body)
	if err != nil {
		return finalRes, header, statusCode, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return finalRes, header, statusCode, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.Client{
		Timeout: 40 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return finalRes, header, statusCode, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return finalRes, header, statusCode, err
	}
	finalRes = respBody
	header = resp.Header
	statusCode = resp.StatusCode
	return finalRes, header, statusCode, err
}
