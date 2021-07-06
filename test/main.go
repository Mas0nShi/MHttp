package main

import (
	"github.com/Mas0nShi/MHttp"
	"github.com/Mas0nShi/goConsole/console"
	"sync"
)

func main() {
	gp := sync.WaitGroup{}

	headers := map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.64",
	}

	for i := 0; i < 100; i++ {
		gp.Add(1)
		go func() {
			http := new(MHttp.MHttp)
			url := "http://127.0.0.1/test.php"
			http.Open("POST", url)
			http.SetRequestHeaders(headers)
			//http.SetProxy("118.251.115.55:11405")
			http.SetCookie("ASP.NET_SessionId", "21231313123")
			http.Send("StudentId=12355566&Name=cxx&acadYears=2020-2021&team=2&type=1")
			console.Log(http.GetResponseText())
			gp.Done()
		}()

	}
	gp.Wait()
}
