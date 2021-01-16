package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Send interface {
	Send(responseBody interface{}) error
}

type Header interface {
	Headers(header map[string]string) ParamsQuerier
}

type RequestQuerier interface {
	ParamsQuerier
	Header
}
type ParamsQuerier interface {
	Params(params []map[string]string) Send
	Send
}

// type Methoder interface {
// 	Get(ctx context.Context, URL string) RequestQuerier
// 	Post(ctx context.Context, URL string, request interface{}) RequestQuerier
// 	Put(ctx context.Context, URL string, request interface{}) RequestQuerier
// 	Delete(ctx context.Context, URL string) RequestQuerier
// }

type RequestStruct struct {
	ctx     context.Context
	request *http.Request
	err     error
}

func (r *RequestStruct) Params(params []map[string]string) Send {

	if r.err != nil {
		return r
	}

	queryParams := r.request.URL.Query()
	for _, param := range params {
		for key, val := range param {
			queryParams.Add(key, val)
		}
	}
	return r

}

func (r *RequestStruct) Headers(headers map[string]string) ParamsQuerier {
	if r.err != nil {
		return r
	}

	for key, val := range headers {
		r.request.Header.Add(key, val)
	}
	return r
}

// func (r *RequestStruct) Delete(ctx context.Context, url string) RequestQuerier {

// 	requst, err := createHttpRequest(url, "DELETE", nil)
// 	r.request = requst
// 	if err != nil {
// 		r.err = err
// 	}
// 	return r
// }

// func Get(ctx context.Context, url string) RequestQuerier {

// 	r := new(RequestStruct)
// 	requst, err := createHttpRequest(url, "GET", nil)
// 	r.request = requst
// 	if err != nil {
// 		r.err = err
// 	}
// 	return r
// }

// func Post(ctx context.Context, url string, body interface{}) RequestQuerier {

// 	r := new(RequestStruct)
// 	requst, err := createHttpRequest(url, "POST", body)
// 	r.request = requst
// 	if err != nil {
// 		r.err = err
// 	}
// 	return r
// }

func Put(ctx context.Context, url string, body interface{}) RequestQuerier {

	r := new(RequestStruct)
	requst, err := createHttpRequest(url, "PUT", body)
	r.request = requst
	if err != nil {
		r.err = err
	}
	return r
}

func Request(ctx context.Context, method, url string, body interface{}) RequestQuerier {

	r := new(RequestStruct)
	requst, err := createHttpRequest(url, method, body)
	r.request = requst
	if err != nil {
		r.err = err
	}
	return r
}

func (r *RequestStruct) Send(body interface{}) error {

	if r.err != nil {
		return r.err
	}

	client := new(http.Client)
	resp, err := client.Do(r.request)
	if err != nil {
		return err
	}
	var buf = new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.Bytes()

	err = json.Unmarshal(responseBody, body)
	if err != nil {
		return err
	}
	return nil
}

func createHttpRequest(url, method string, requestBody interface{}) (*http.Request, error) {

	bodytoBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(bodytoBytes)
	request, err := http.NewRequest(method, url, r)
	if err != nil {
		return nil, err
	}
	return request, nil

}
