package main

import (
	"crypto/tls"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"yannscrapy/logging"
)

func main() {
	GetMyPage()
}

func GetMyPage() {
	url := "https://www.douban.com/people/219914315/"
	//url := "https://news.baidu.com"

	header := map[string]string{
		"Accept":       "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	//	"Accept-Encoding": "gzip, deflate, br",

		"Accept-Charset": "utf-8",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Cache-Control":   "max-age=0",
		"Connection":      "keep-alive",
		"Cookie":          `douban-fav-remind=1; __utmc=30149280; gr_user_id=91212bc2-ab6d-43f3-a369-27bf206362c0; _vwo_uuid_v2=D3CDB75E7EA9CF0EF87F9CC6F186B8813|8e96b58e4f33c0b871f09799a9387637; apiKey=; douban-profile-remind=1; __utmv=30149280.21991; bid=ud28GnQ7tuw; __yadk_uid=jpL5rNl59sRY7lBaWJdwuCFUXF2ygpjQ; viewed="17584179_25809330_4292217_17557980_21323198"; ll="118281"; push_noty_num=0; push_doumail_num=0; __gads=ID=67aa2008465f1c8d:T=1605342393:S=ALNI_MagFH3npfNchMqhXwNOd6LX4tUg_g; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1628921275%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3Dp824eYFDDx9fkBTaA_qpXZK3A3tIW-wUGo9eje0_9SCtBsnJKcO18teizOQHiSOl%26wd%3D%26eqid%3Dbbab457300002f210000000661175db6%22%5D; _pk_ses.100001.8cb4=*; ap_v=0,6.0; __utma=30149280.2007004949.1595730645.1628870782.1628921276.54; __utmz=30149280.1628921276.54.34.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmt=1; dbcl2="219914315:IrR332PCbzs"; ck=gi_j; _pk_id.100001.8cb4=c0cb64199260533e.1595670983.36.1628921513.1628872681.; __utmb=30149280.17.10.1628921276`,
		"Host":            "www.douban.com",
		"Referer":         "https://www.douban.com/people/219914315/",
		//sec-ch-ua: " Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"
		"sec-ch-ua-mobile":          "?0",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
	}


	//request := resty.R()
	//client := resty.New().SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true }).SetTimeout(60)
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	//client.GetClient().CheckRedirect = func(req *http.Request, via []*http.Request) error {
	//	return http.ErrUseLastResponse
	//}
	request := client.R()
	SetHeader(request, header)
	resp, err := request.Get(url)
	if err != nil {
		logging.Panic(err)
	}

	logging.Infof("statusCode=%v", resp.StatusCode())
	logging.Infof("status=%v", resp.Status())

	logging.Infof("------header-----")
	for e, v := range resp.Header() {
		logging.Infof("%v:%v", e, v)
	}

	dstFileName := "mydouban.html"
	ioutil.WriteFile(dstFileName, resp.Body(), 0600)

	//logging.Infof("body=%s", string(resp.Body()))

	//utf8Str := mahonia.NewDecoder("gbk").ConvertString(string(resp.Body()))
	//logging.Infof("body=%s", utf8Str)

}
func SetHeader(request *resty.Request, m map[string]string) {
	for k, v := range m {
		request.SetHeader(k, v)
	}
}
