package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/resty.v1"
	//"yannscrapy/logging"
)

var header1 = map[string]string{

	"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-language": "zh-CN,zh;q=0.9",
	"cache-control":   "max-age=0",

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
	"authority": "www.anadf.com",
	"method":    "POST",
	"path":      "/cn/MemberLogin.aspx",
	"scheme":    "https",

	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-encoding":           "gzip, deflate, br",
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

func main() {
	client, err := CreateClient()
	if err != nil {
		fmt.Println(err)
	}
	url := "https://www.anadf.com/cn/MemberLogin.aspx"

	request := client.R()
	rsp1, _ := GetRequest(request, url, header1, "MemberLogin.html")
	fmt.Printf("---------------------------------------\n")
	fmt.Printf("---------------------------------------\n")

	fmt.Printf(" get response cookies:\n")
	for i, cookie := range rsp1.Cookies() {
		fmt.Printf("%v,%v\n", i, cookie.String())
	}
	fmt.Printf("get response cookies end\n")

	request = client.R()
	copyRequestParam(rsp1, request)

	cookieStr := CopyCookies(rsp1)
	request.SetHeader("cookie", cookieStr)
	PostRequest(request, url, header, nil, "login_result.html")

}

func copyRequestParam(response *resty.Response, request *resty.Request) {
	//dom, err := goquery.NewDocumentFromReader(rsp1.RawBody())
	//fmt.Printf("body=%v", string(rsp1.Body()))
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(response.Body())))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	__EVENTTARGET, _ := dom.Find("input#__EVENTTARGET").Eq(0).Attr("value")
	fmt.Printf("__EVENTTARGET=%v\n", __EVENTTARGET)

	__EVENTARGUMENT, _ := dom.Find("input#__EVENTARGUMENT").Eq(0).Attr("value")
	fmt.Printf("__EVENTARGUMENT=%v\n", __EVENTARGUMENT)

	__LASTFOCUS, _ := dom.Find("input#__LASTFOCUS").Eq(0).Attr("value")
	fmt.Printf("__LASTFOCUS=%v\n", __LASTFOCUS)

	__VIEWSTATE, _ := dom.Find("input#__VIEWSTATE").Eq(0).Attr("value")
	fmt.Printf("__VIEWSTATE=%v\n", __VIEWSTATE)

	__VIEWSTATEGENERATOR, _ := dom.Find("div input#__VIEWSTATEGENERATOR").Eq(0).Attr("value")
	fmt.Printf("__VIEWSTATEGENERATOR=%v\n", __VIEWSTATEGENERATOR)

	__EVENTVALIDATION, _ := dom.Find("div input#__EVENTVALIDATION").Eq(0).Attr("value")
	fmt.Printf("__EVENTVALIDATION=%v\n", __EVENTVALIDATION)

	ddlAirport, _ := dom.Find("div select[name='ctl00$ddlAirport'] option[selected=selected]").Eq(0).Attr("value")
	fmt.Printf("ctl00$ddlAirport=%v\n", ddlAirport)

	ddlLanguage, _ := dom.Find("div select[name='ctl00$ddlLanguage'] option[selected=selected]").Eq(0).Attr("value")
	fmt.Printf("ctl00$ddlLanguage=%v\n", ddlLanguage)

	txtKeyword := dom.Find("div input[name='ctl00$txtKeyword']").Eq(0).Text()
	fmt.Printf("ctl00$txtKeyword=%v\n", txtKeyword)

	txtMail := dom.Find("div input[name='ctl00$ContentPlaceHolder1$txtMail']").Eq(0).Text()
	fmt.Printf("ctl00$ContentPlaceHolder1$txtMail=%v\n", txtMail)

	txtPass := dom.Find("div input[name='ctl00$ContentPlaceHolder1$TxtPASS']").Eq(0).Text()
	fmt.Printf("ctl00$ContentPlaceHolder1$TxtPASS=%v\n", txtPass)

	btnLogin, _ := dom.Find("div input[name='ctl00$ContentPlaceHolder1$btnLogin']").Eq(0).Attr("value")
	fmt.Printf("ctl00$ContentPlaceHolder1$btnLogin=%v\n", btnLogin)

	form := map[string]string{
		"__EVENTTARGET":        __EVENTTARGET,
		"__EVENTARGUMENT":      __EVENTARGUMENT,
		"__LASTFOCUS":          __LASTFOCUS,
		"__VIEWSTATEGENERATOR": __VIEWSTATEGENERATOR,
		"__VIEWSTATE":          __VIEWSTATE,
		"__EVENTVALIDATION":    __EVENTVALIDATION,

		"ctl00$ddlAirport":                  ddlAirport,
		"ctl00$ddlLanguage":                 ddlLanguage,
		"ctl00$txtKeyword":                  txtKeyword,
		"ctl00$ContentPlaceHolder1$txtMail": "getway@moran.cn", //   "sdsdw@126.com"
		"ctl00$ContentPlaceHolder1$TxtPASS": "moranjiuye1",     // "123123ab"
		// "ctl00$ContentPlaceHolder1$txtMail":  "sdsdw@126.com", //"getway@moran.cn",    //"getway@moran.cn"
		// "ctl00$ContentPlaceHolder1$TxtPASS":  "123123ab",      //"moranjiuye1",
		"ctl00$ContentPlaceHolder1$btnLogin": btnLogin,
	}

	request.SetFormData(form)
}

func PostRequest(request *resty.Request,
	url string,
	header map[string]string,
	form map[string]string,
	saveFile string) (*resty.Response, error) {

	if header != nil {
		SetHeader(request, header)
	}

	if form != nil {
		request.SetFormData(form)
	}

	fmt.Printf("--------------------------------------\n")
	fmt.Printf("--------------------------------------\n")
	fmt.Printf("post request form:%v\n", request.FormData.Encode())
	//fmt.Printf("reqeust:%v\n",request.Body)
	fmt.Printf("post request header:%v\n", request.Header)

	resp, err := request.Post(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("--------------------------------------\n")
	fmt.Printf("status=%v\n", resp.Status())
	fmt.Printf("------post response header-----\n")
	for e, v := range resp.Header() {
		fmt.Printf("%v:%v\n", e, v)
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
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.GetClient().Jar = jar
	return client, nil
}

func GetRequest(request *resty.Request, url string, header map[string]string, saveFile string) (*resty.Response, error) {
	if header != nil {
		SetHeader(request, header)
	}
	resp, err := request.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("--------------------------------------\n")
	fmt.Printf("status=%v\n", resp.Status())
	fmt.Printf("------get response header-----\n")
	for e, v := range resp.Header() {
		fmt.Printf("%v:%v\n", e, v)
	}

	ioutil.WriteFile(saveFile, resp.Body(), 0600)
	return resp, nil
}

func CopyCookies(response *resty.Response) string {
	var cookieStr string

	cookies := response.Cookies()
	for _, c := range cookies {
		if cookieStr == "" {
			cookieStr = c.Name + "=" + c.Value
		} else {
			cookieStr = cookieStr + ";" + c.Name + "=" + c.Value
		}
	}

	return cookieStr
}
