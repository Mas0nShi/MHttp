# MHttp - easyHTTP in Go

## Feature
1. 🚀 Easy to use 🚀
2. 🐲 object-oriented 🐲
3. ✈️Depend on `net/http` ✈️

## Usage
1. install MHttp

`go get github.com/Mas0nShi/MHttp`

2. enjoy
```go

package main

import (
	"fmt"
	"github.com/Mas0nShi/MHttp"
)

func main() {
	headers := map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.64",
	}
	http := new(MHttp.MHttp)
	url := "http://127.0.0.1/test.php"
	http.Open("POST", url)
	http.SetRequestHeaders(headers)
	http.SetCookie("cool", "123456")
	http.Send("cardId=123456&Name=Mas0n&type=1")
	fmt.Println(http.GetResponseText())
}
```

## Encapsulation
    Open(method string, url string)
    Send(body interface{})
    SetCookie(key string, value string)
    SetCookies(cookies map[string]string)
    SetRequestHeader(key string, value string)
    SetRequestHeaders(headers map[string]string)
    SetProxy(url string)
    GetHttpCode() int 
    GetResponseBody() []byte
    GetResponseText() string
    GetResponseHeader(key string) []string
    GetResponseHeaders() map[string][]string
    GetCookie(key string) string
    GetCookies() string
    Clear() 
    AutoHeaders()

## TODO
-[ ] Map read/write conflict during concurrency
