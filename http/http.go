package http

// https://i6448038.github.io/2017/11/11/httpAndGolang/
// https://www.cnblogs.com/zhaof/p/11346412.html

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	
	p "github.com/syyongx/php2go"
)

// postData url.Values{"key": {"Value"}, "id": {"123"}}
func PostForm(postUrl string, timeout int, postData url.Values, headers map[string]interface{}) (string, int, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	
	resp, err := client.PostForm(postUrl, postData)
	if err != nil {
		return "", 0, err
	}
	
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	
	return string(body), resp.StatusCode, err
}

func PostFormComplex(postUrl string, timeout int, postData url.Values, headers map[string]interface{}) (string, int, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	request, err := http.NewRequest(http.MethodPost, postUrl, strings.NewReader(postData.Encode()))
	if err != nil {
		return "", 0, err
	}
	
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Content-Length", strconv.Itoa(len(postData.Encode())))
	if !p.Empty(headers) {
		for hk, hv := range headers {
			request.Header.Set(hk, hv.(string))
			// request.Header.Set(hk, tool.InterfaceToString(hv))
		}
	}
	// todo cookie
	cookie := &http.Cookie{Name: "sso_token_yundunv5", Value: "cugr0cwhgpsggkggs80040okc", HttpOnly: true}
	request.AddCookie(cookie)
	
	resp, err := client.Do(request)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		return "", 0, err
	}
	
	return string(body), 0, err
}

func SendJson(url string, method string, data interface{}, timeout int, headers map[string]interface{}) (string, int, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", 0, err
	}
	reader := bytes.NewReader(jsonData)
	request, err := http.NewRequest(method, url, reader)
	if err != nil {
		return "", 0, err
	}
	request.Header.Set("Content-Type", "application/json")
	if !p.Empty(headers) {
		for hk, hv := range headers {
			request.Header.Set(hk, hv.(string))
		}
	}
	// todo cookie
	cookie := &http.Cookie{Name: "sso_token_yundunv5", Value: "cugr0cwhgpsggkggs80040okc", HttpOnly: true}
	request.AddCookie(cookie)
	
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		return "", 0, err
	}
	
	return string(body), 0, err
}

func Get(url string, timeout int, headers map[string]interface{}) (string, int, error) {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}
	
	if !p.Empty(headers) {
		for hk, hv := range headers {
			request.Header.Set(hk, hv.(string))
		}
	}
	// todo cookie
	cookie := &http.Cookie{Name: "sso_token_yundunv5", Value: "cugr0cwhgpsggkggs80040okc", HttpOnly: true}
	request.AddCookie(cookie)
	
	resp, err := client.Do(request)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	
	if err != nil {
		return "", 0, err
	}
	
	return string(body), 0, err
}

func Send(method string, url string, timeout int, headers map[string]interface{}, reqParams interface{}) (body string, code int, err error) {
	m := strings.ToUpper(method)
	switch m {
	case "GET":
		return Get(url, timeout, headers)
	case "POST":
		return SendJson(url, http.MethodPost, reqParams, timeout, headers)
	case "PUT":
		return SendJson(url, http.MethodPut, reqParams, timeout, headers)
	case "DELETE":
		return SendJson(url, http.MethodDelete, reqParams, timeout, headers)
	}
	
	return "UNKNOW", 0, nil
}

type JsonRequest struct {
	Url     string
	Method  string
	Body    []byte
	Timeout int
	Header  map[string]string
	Cookie  *http.Cookie
}

func JsonReq(r *JsonRequest) (string, int, error) {
	reader := bytes.NewReader(r.Body)
	request, err := http.NewRequest(r.Method, r.Url, reader)
	if err != nil {
		return "", 0, err
	}
	
	request.Header.Set("Content-Type", "application/json")
	if !p.Empty(r.Header) {
		for k, v := range r.Header {
			request.Header.Set(k, v)
		}
	}
	
	request.AddCookie(r.Cookie)
	
	client := &http.Client{
		Timeout: time.Duration(r.Timeout) * time.Second,
	}
	
	resp, err := client.Do(request)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}
	
	return string(body), 0, err
}
