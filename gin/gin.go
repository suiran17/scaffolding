package gin

import (
	"io"
	"net/http"
	"net/http/httptest"
	
	"github.com/gin-gonic/gin"
	strings "github.com/suiran17/scaffolding/string"
)

const (
	HTTPMethodGET    = http.MethodGet
	HTTPMethodPOST   = http.MethodPost
	HTTPMethodPUT    = http.MethodPut
	HTTPMethodDELETE = http.MethodDelete
)

type Request struct {
	R      *gin.Engine
	Url    string
	Herder map[string]string
	Method string
	Body   io.Reader
}

type Response struct {
	StatusCode int
	Body       []byte
	Response   *http.Response
}

func Req(request *Request) (*Response, string, error) {
	req, err := http.NewRequest(request.Method, request.Url, request.Body)
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
