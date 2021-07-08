package MHttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	url2 "net/url"
	"strings"
)

type MHttp struct {
	method string
	url    string
	res    response
	req    request
}
type request struct {
	body    io.Reader
	headers map[string]string
	cookies map[string]string
	proxy   *http.Transport
}
type response struct {
	body     []byte
	httpCode int
	headers  http.Header
	cookies  []*http.Cookie
}

func (h *MHttp) GetHttpCode() int {
	return h.res.httpCode
}
func (h *MHttp) GetResponseBody() []byte {
	return h.res.body
}
func (h *MHttp) GetResponseText() string {
	return Bytes2str(h.res.body)
}
func (h *MHttp) GetResponseHeader(key string) []string {
	return h.res.headers[key]
}
func (h *MHttp) GetResponseHeaders() map[string][]string {
	return h.res.headers
}
func (h *MHttp) GetCookie(key string) string {
	for _, cookie := range h.res.cookies {
		if cookie.Name == key {
			return cookie.String()
		}
	}
	return ""
}
func (h *MHttp) GetCookies() string {
	cookies := ""
	for _, cookie := range h.res.cookies {
		cookies += cookie.String()
	}
	return cookies
}

func (h *MHttp) SetCookie(key string, value string) {
	h.req.cookies[key] = value

}
func (h *MHttp) SetCookies(cookies map[string]string) {
	for k, s := range cookies {
		h.req.cookies[k] = s
	}
}
func (h *MHttp) SetRequestHeader(key string, value string) {
	h.req.headers[key] = value
}
func (h *MHttp) SetRequestHeaders(headers map[string]string) {
	h.req.headers = headers
}
func (h *MHttp) SetProxy(ip string) {
	if ip != "" {
		if !strings.HasPrefix(ip, "http://") {
			ip = "http://" + ip
		}
		parse, err := url2.Parse(ip)
		if err != nil {
			panic("MHttp/SetProxy error in parse url.")
		}
		h.req.proxy = &http.Transport{Proxy: http.ProxyURL(parse)}
	} else {
		h.req.proxy = &http.Transport{}
	}
}

func (h *MHttp) Clear() {
	h.req = *new(request)
	h.res = *new(response)
	h.method = ""
	h.url = ""
}
func (h *MHttp) Open(method string, url string) {
	h.url = url
	h.method = method
	h.req.proxy = &http.Transport{}
	if h.req.cookies == nil {
		h.req.cookies = map[string]string{}
	}

	if h.req.headers == nil {
		h.req.headers = map[string]string{}
	}
	if h.req.headers["Accept"] == "" {
		h.req.headers["Accept"] = "*/*"
	}
	if h.req.headers["Accept-Language"] == "" {
		h.req.headers["Accept-Language"] = "zh-cn"
	}
	if h.req.headers["Referer"] == "" {
		h.req.headers["Referer"] = h.url
	}
	if h.method == "POST" && h.req.headers["Content-Type"] == ""{
		h.req.headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

}
func (h *MHttp) Send(body interface{}) {
	switch v := body.(type) {
	case nil:
		break
	case string:
		h.req.body = strings.NewReader(v)
	case []byte:
		h.req.body = bytes.NewReader(v)
	default:
		panic("body type error.")
	}

	req, err := http.NewRequest(h.method, h.url, h.req.body)
	if err != nil {
		panic(err)
	}

	// set headers
	for k, v := range h.req.headers {
		req.Header.Add(k, v)
	}

	// set cookies
	jar, _ := cookiejar.New(nil)
	icookies := make([]*http.Cookie,len(h.req.cookies))
	if len(h.req.cookies) > 0 {
		cout := 0
		for key, value := range h.req.cookies {
			icookies[cout] = &http.Cookie{Name: key,Value: value,HttpOnly: true}
			cout++
		}
	}
	jar.SetCookies(req.URL, icookies)
	// send http requests
	client := &http.Client{Transport: h.req.proxy}
	client.Jar = jar
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// read response
	resp, err := ioutil.ReadAll(res.Body)
	h.res = response{resp, res.StatusCode, res.Header, res.Cookies()}
}
