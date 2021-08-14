package main

import (
	"crypto/tls"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"yannscrapy/logging"
)


var header1 = map[string]string{

	"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language": "zh-CN,zh;q=0.9",
	"cache-control": "max-age=0",

	"sec-ch-ua":                 `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
}

var header = map[string]string{
	"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	//"accept-encoding": "gzip, deflate, br",
	"accept-language":           "zh-CN,zh;q=0.9",
	"cache-control":             "max-age=0",
	"content-type":              "application/x-www-form-urlencoded",
	"origin":                    "https://www.anadf.com",
	"referer":                   "https://www.anadf.com/cn/MemberLogin.aspx",
	"sec-ch-ua":                 `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
}



var formMap = map[string]string{
	"__EVENTTARGET":        "",
	"__EVENTARGUMENT":      "",
	"__LASTFOCUS":          "",
	"__VIEWSTATE":          "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",
	"__VIEWSTATEGENERATOR": "AA4F3242",
	"__EVENTVALIDATION":    "ZDk2ZsodNY7TzoAd8ufH9ESxyRZrKeSnrOH3mjuFZXZwGyJSJ85C8XPN4Noebg/qsqDPVmVE31pbUlmf5UkSDbTbMFJEZ9UNOHpbCBx7UDgwvhk5L0ZZmvQsw27fmb8lghyZBgLUxgJe1OQ7oQrNIWPezg/IDi5yY5FfwgJ3bCZdhZTDw3JJSGQEBItDn3OphTaPuvuHo2JZqvbEkuxW8687n0zpir41D5aESo2Wq/3hAr9FuWgfJBy6KQSZstACJKJ66jk+t3WcqoQdgbUn6ONgu5rHFvlSR4e9HVynFcVNsdDvHSnNwsz5wuecAL0p3Ikhi0llhAj1oDQjlqKG+/lRPhK545FkTnmpWhgnPtIPynSDVcv+GQ6JXZSYiZOXiDrNhyb6u4iq4aS8isX/NAburPxF/G5TLnSLRR4UdREqBhdPYDYBZn2pJK8J9yp25G8YQZOVByXx/slOTKIuggkSuCxDDVitYoaTpn8UB7Cpn1A+G8IWJZZkInridCB5NF7APFGcl8fRqeMQm3sJBnwvZKw=",

	"ctl00$ddlAirport":                   "04",
	"ctl00$ddlLanguage":                  "2",
	"ctl00$txtKeyword":                   "",
	"ctl00$ContentPlaceHolder1$txtMail":  "getway@moran.cn",
	"ctl00$ContentPlaceHolder1$TxtPASS":  "moranjiuye1",
	"ctl00$ContentPlaceHolder1$btnLogin": "登录",
}

func main() {
	client, err := CreateClient()
	if err != nil {
		logging.Panic(err)
	}
	request := client.R()
	url := "https://www.anadf.com/cn/MemberLogin.aspx"
	GetRequest(request, url, header1, "MemberLogin.html")

	logging.Infof("---------------------------------------")
	logging.Infof("---------------------------------------")

	//url = "https://www.anadf.com/cn/MemberLogin.aspx"

	request = client.R()
	PostRequest(request, url, header, formMap, "login.html")
}

func PostRequest(request *resty.Request,
		url string,
		header map[string]string,
		form map[string]string,
		saveFile string) (*resty.Response, error) {

	SetHeader(request, header)
	request.SetFormData(form)

	resp, err := request.Post(url)
	if err != nil {
		logging.Panic(err)
	}

	logging.Infof("status=%v", resp.Status())
	logging.Infof("------header-----")
	for e, v := range resp.Header() {
		logging.Infof("%v:%v", e, v)
	}

	ioutil.WriteFile(saveFile, resp.Body(), 0600)
	return resp, nil
}

func SetHeader(request *resty.Request, m map[string]string) {
	if m == nil {
		return
	}
	for k, v := range m {
		request.SetHeader(k, v)
	}
}

func CreateClient() (*resty.Client, error) {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.GetClient().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return client, nil
}


func CreateRequest() (*resty.Request, error) {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.GetClient().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	request := client.R()
	return request, nil
}

func GetRequest(request *resty.Request, url string, header map[string]string, saveFile string) (*resty.Response, error){
	SetHeader(request, header)
	resp, err := request.Get(url)
	if err != nil {
		logging.Panic(err)
	}
	logging.Infof("status=%v", resp.Status())
	logging.Infof("------header-----")
	for e, v := range resp.Header() {
		logging.Infof("%v:%v", e, v)
	}

	ioutil.WriteFile(saveFile, resp.Body(), 0600)
	return resp, nil
}
