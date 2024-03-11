package httpclient

func (hc *HttpClient) Get(url string, payload map[string]string) ([]byte, int, error) {

	// create request
	request := hc.client.R()

	// set request payload if provided
	if len(payload) > 0 {
		request = request.SetQueryParams(payload)
	}

	// default client headers
	request = request.SetHeaders(defaultHeaders)

	// call get request
	resp, err := request.Get(url)
	if err != nil {
		return nil, 0, err
	}

	return resp.Body(), resp.StatusCode(), nil
}

func (hc *HttpClient) Post(url string, payload map[string]string) ([]byte, int, error) {

	// create request
	request := hc.client.R()

	// set request payload if provided
	if len(payload) > 0 {
		request = request.SetBody(payload)
	}

	// default client headers
	request = request.SetHeaders(defaultHeaders)

	// call post request
	resp, err := request.Post(url)
	if err != nil {
		return nil, 0, err
	}

	return resp.Body(), resp.StatusCode(), nil
}
