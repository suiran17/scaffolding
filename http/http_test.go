package http

import (
	"net/http"
	"testing"
)

func TestJsonRequest(t *testing.T) {
	header := map[string]string{}
	url := ""
	data := []byte{}
	timeout := 30
	// cookie := &http.Cookie{Name: "sso_token_yundunv5", Value: "cugr0cwhgpsggkggs80040okc", HttpOnly: true}
	cookie := &http.Cookie{}
	
	r := &JsonRequest{
		Url:     url,
		Method:  http.MethodPost,
		Body:    data,
		Timeout: timeout,
		Header:  header,
		Cookie:  cookie,
	}
	body, code, err := JsonReq(r)
	t.Log(body, code, err)
}
