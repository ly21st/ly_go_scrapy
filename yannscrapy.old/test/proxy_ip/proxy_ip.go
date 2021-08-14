package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

var total, useful int = 0, 0
var status = make(chan int)

func main() {

	xcurl := "http://www.xicidaili.com/wt/"
	request, _ := http.NewRequest("GET", xcurl, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")

	cli1 := &http.Client{}
	response, err := cli1.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	dom, _ := goquery.NewDocumentFromResponse(response)

	dom.Find("#ip_list tbody tr").Each(func(i int, context *goquery.Selection) {
		ip := context.Find("td").Eq(1).Text()
		port := context.Find("td").Eq(2).Text()
		httpType := context.Find("td").Eq(5).Text()
		proxyIp := strings.ToLower(httpType) + "://" + ip + ":" + port
		// nim := context.Find("td").Eq(4).Text() //是否是高匿,高匿的可以隐藏你的原始IP

		if ip != "" && port != "" {
			total++
			go checkProxyIP(proxyIp, i)
		}
	})

	for i := 0; i < total; i++ {
		<-status
	}
	fmt.Println("num=", total, "\nuseful=", useful)
	fmt.Println("END!")
}

func checkProxyIP(proxyIp string, i int) {
	req, _ := http.NewRequest("GET", "http://test.bestbing.cn/", nil) //这里自己搭个web服务验证代理是否可用
	proxy, _ := url.Parse(proxyIp)
	cli2 := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}
	resp, _ := cli2.Do(req)

	if resp != nil && resp.StatusCode == 200 {
		buf := make([]byte, 1024)
		n, _ := resp.Body.Read(buf)
		reStr := string(buf[:n])

		if reStr == "Hello World" { //验证代理有没有做手脚，可能给你返回一堆广告
			useful++
			fmt.Println(proxyIp)
		}

	}
	status <- i
}