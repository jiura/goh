package goh

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

/*
Sends an HTTP request. Returns a pointer to the response, its status code as an int and its body as a string.

Can't add the same header more than once; the one that comes last will be used to set the value for that key.

Notice that the body from the http.Response will be closed when this function returns, so the body should only be accessed throught the returned string.
*/
func HttpRequest(method, url string, body []byte, headers map[string]string) (*http.Response, int, string, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, 0, "", err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, "", err
	}
	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, resp.StatusCode, "", err
	}

	return resp, resp.StatusCode, string(resp_body), nil
}

/*
Sends an HTTP request. Returns a pointer to the response, its status code as an int and its body as a map[string]any.

A JSON body is expected in the response.

Can't add the same header more than once; the one that comes last will be used to set the value for that key.

Notice that the body from the http.Response will be closed when this function returns, so the body should only be accessed throught the returned map[string]string.
*/
func HttpRequestJson(method, url string, body []byte, headers map[string]string) (*http.Response, int, map[string]any, error) {
	resp, status_code, resp_body, err := HttpRequest(method, url, body, headers)
	if err != nil {
		return resp, status_code, nil, err
	}

	var resp_body_json map[string]any
	if err = json.Unmarshal([]byte(resp_body), &resp_body_json); err != nil {
		return resp, status_code, nil, errors.New(err.Error() + " - Response body: " + resp_body)
	}

	return resp, status_code, resp_body_json, nil
}
