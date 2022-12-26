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

var HeaderJson = map[string]string{"Content-Type": "application/json;charset=UTF-8"}

type Request struct {
	R      *gin.Engine
	Url    string
	Herder map[string]string
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
	
	for k, v := range request.Herder {
		req.Header.Set(k, v)
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
