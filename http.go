package goh

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

/* TYPES START */

type HttpResponse struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

type HttpResponseJson struct {
	HttpResponse
	Body Json
}

/* TYPES END */

/* FUNCTIONS START */

/*
Sends an HTTP request. Returns a pointer to a response struct with status, headers and body as []byte.

Can't add the same header more than once; the one that comes last will be used to set the value for that key.

Notice that the body from the http.Response will be closed when this function returns, so the body should only be accessed through the returned string.
*/
func HttpRequest(method, url string, headers map[string]string, body []byte) (*HttpResponse, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &HttpResponse{resp.Status, resp.StatusCode, resp.Header, nil}, err
	}

	return &HttpResponse{resp.Status, resp.StatusCode, resp.Header, resp_body}, nil
}

/*
Sends an HTTP request. Returns a pointer to a response struct, with status, headers and body as a Json (map[string]any).

A JSON body is expected in the response.

Can't add the same header more than once; the one that comes last will be used to set the value for that key.

Notice that the body from the http.Response will be closed when this function returns, so the body should only be accessed through the returned Json.
*/
func HttpRequestJson(method, url string, headers map[string]string, body []byte) (*HttpResponseJson, error) {
	resp, err := HttpRequest(method, url, headers, body)
	if err != nil {
		return &HttpResponseJson{*resp, nil}, err
	}

	var resp_body_json Json
	if err = json.Unmarshal(resp.Body, &resp_body_json); err != nil {
		return &HttpResponseJson{*resp, nil}, errors.New(err.Error() + " - Response body: " + string(resp.Body))
	}

	return &HttpResponseJson{*resp, resp_body_json}, nil
}

/* FUNCTIONS END */
