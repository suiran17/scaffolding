package gin

import (
	"io"
	"net/http"
	"net/http/httptest"
	
	"github.com/gin-gonic/gin"
	strings "github.com/suiran17/scaffolding/string"
)

const (
	HTTPMethodGET    = "GET"
	HTTPMethodPOST   = "POST"
	HTTPMethodPUT    = "PUT"
	HTTPMethodDELETE = "DELETE"
)

var HeaderJson = &Header{Key: "Content-Type", Value: "application/json;charset=UTF-8"}

type Header struct {
	Key   string
	Value string
}

type HeaderSet []*Header

type Request struct {
	R      *gin.Engine
	Url    string
	Herder HeaderSet
	Method string
	Reader io.Reader
}

type Response struct {
	StatusCode int
	Body       []byte
	Response   *http.Response
}

func Req(request Request) (*Response, string, error) {
	req, err := http.NewRequest(request.Method, request.Url, request.Reader)
	if err != nil {
		return &Response{}, "", err
	}
	
	for _, h := range request.Herder {
		req.Header.Set(h.Key, h.Value)
	}
	
	rec := httptest.NewRecorder()
	
	request.R.ServeHTTP(rec, req)
	
	response := rec.Result()
	
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return &Response{}, "", err
	}
	defer response.Body.Close()
	
	return &Response{
		StatusCode: response.StatusCode,
		Body:       body,
		Response:   response,
	}, strings.JsonIndent(body), nil
}
