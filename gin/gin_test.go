package gin

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	// "net/url"
	
	"github.com/gin-gonic/gin"
)

// 构建一个简单的 gin 服务

// func TestGinSrv(t *testing.T) {
//
// }

func Handler(c *gin.Context) {
	a := c.PostForm("a")
	b := c.PostForm("b")
	ia, err1 := strconv.Atoi(a)
	ib, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bad params"})
		return
	}
	result := ia + ib
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	
	groupTest := r.Group("/v1")
	{
		groupTest.POST("/test", Handler)
	}
	return r
}

// curl --location --request POST 'localhost:8080/v1/test' \
// --form 'a="1"' \
// --form 'b="2"'
func TestRunGin(t *testing.T) {
	r := setupRouter()
	if err := r.Run(); err != nil {
		return
	}
}

func TestGinExample(t *testing.T) {
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded", // 必须, https://www.cnblogs.com/mafeng/p/7068837.html
		// "Content-Type": "multipart/form-data;charset=UTF-8",
		// "Content-Type": "multipart/form-data; boundary=<calculated when request is sent>",
	}
	
	data := url.Values{
		"a": {"1"},
		"b": {"2"},
	}
	
	request := Request{
		R:      setupRouter(),
		Url:    "/v1/test",
		Herder: header,
		Method: HTTPMethodPOST,
		Reader: bytes.NewBufferString(data.Encode()),
	}
	
	resp, stringBody, err := Req(request)
	if err != nil {
		log.Println("Req err: ", err)
	}
	defer resp.Response.Body.Close()
	
	fmt.Println(resp.StatusCode, string(resp.Body), resp.Response, stringBody)
	
}
