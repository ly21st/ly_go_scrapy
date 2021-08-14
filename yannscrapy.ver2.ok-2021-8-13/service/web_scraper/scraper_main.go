package web_scraper

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
	"yannscrapy/logger"
	"yannscrapy/logging"
	"yannscrapy/utils"
)



func Scrapy_main() {
	urlstr := "https://news.baidu.com"
	// urlstr := "https://www.anadf.com/"
	u, err := url.Parse(urlstr)
	if err != nil {
		logger.Fatal(err)
	}
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.SetProxyFunc(randomProxySwitcher)

	// 超时设定
	c.SetRequestTimeout(100 * time.Second)
	// 指定Agent信息
	// c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36"
	c.OnRequest(func(r *colly.Request) {
		// Request头部设定
		r.Headers.Set("Host", u.Host)
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Origin", u.Host)
		r.Headers.Set("Referer", urlstr)
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN, zh;q=0.9")

		logger.Infof("visiting %v", r.URL)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		// log.Println("title:", e.Text)
		doc := e.DOM
		logger.Infof("title:%v", doc.Text())
	})

	c.OnResponse(func(resp *colly.Response) {
		logger.Infof("response received %v", resp.StatusCode)

		// goquery直接读取resp.Body的内容
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body))

		// 读取url再传给goquery，访问url读取内容，此处不建议使用
		// htmlDoc, err := goquery.NewDocument(resp.Request.URL.String())

		if err != nil {
			logger.Fatal(err)
		}

		// 找到抓取项 <div class="hotnews" alog-group="focustop-hotnews"> 下所有的a解析
		htmlDoc.Find(".hotnews a").Each(func(i int, s *goquery.Selection) {
			band, _ := s.Attr("href")
			title := s.Text()
			logger.Infof("热点新闻 %d: %s - %s", i, title, band)
			// c.Visit(band)
		})

	})

	c.OnError(func(resp *colly.Response, errHttp error) {
		err = errHttp
		logger.Infof("err=%v", err)
	})

	err = c.Visit(urlstr)
}

func randomProxySwitcher(_ *http.Request) (*url.URL, error) {
	// return proxies[random.Intn(len(proxies))], nil
	//var proxies []*url.URL = []*url.URL{
	//	&url.URL{Host: "127.0.0.1:8080"},
	//	&url.URL{Host: "127.0.0.1:8081"},
	//}

	//var proxies []*url.URL = []*url.URL{
	//	&url.URL{Host: "175.143.37.162:80"},
	//}

	var proxies []*url.URL = make([]*url.URL, 0)
	ipList, err := utils.GetDefaultAllHealthProxyIps()
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	for _, ip := range ipList {
		var proxyUrl *url.URL
		logger.Infof("ip=%s", ip)
		if strings.HasPrefix(ip,"http://") {
			proxyUrl = &url.URL{Host: ip[7:]}
		} else if strings.HasPrefix(ip, "https://") {
			proxyUrl = &url.URL{Host: ip[8:]}
		} else {
			proxyUrl = &url.URL{Host: ip}
		}
		proxies = append(proxies, proxyUrl)
	}

	logger.Infof("proxies=%v", proxies)
	return proxies[rand.Intn(len(proxies))], nil
}