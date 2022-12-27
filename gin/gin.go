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

func New() *Request {
	return &Request{}
}

func (r *Request) SetEngine(e *gin.Engine) *Request {
	r.R = e
	return r
}

func (r *Request) SetUrl(url string) *Request {
	r.Url = url
	return r
}

func (r *Request) SetHeader(header map[string]string) *Request {
	r.Herder = header
	return r
}

func (r *Request) SeMethod(method string) *Request {
	r.Method = method
	return r
}

func (r *Request) SetBody(body io.Reader) *Request {
	r.Body = body
	return r
}

func (r *Request) Req() (*Response, string, error) {
	req, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return &Response{}, "", err
	}
	
	for k, v := range r.Herder {
		req.Header.Set(k, v)
	}
	
	rec := httptest.NewRecorder()
	
	r.R.ServeHTTP(rec, req)
	
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
