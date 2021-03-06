package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/resty.v1"
	//"yannscrapy/logging"
)

func main() {
	beginTime := time.Now()
	fmt.Printf("begin time:%v\n", beginTime.String())

	client, err := CreateClient()
	if err != nil {
		fmt.Println(err)
	}
	//url := "https://www.anadf.com/cn/MemberLogin.aspx"
	fmt.Printf("-----search item-----")
	url := "https://www.anadf.com/cn/ItemDetail.aspx?S_CD=4020102654"
	request := client.R()
	header := CommonGetHeader()
	rsp, _ := GetRequest(request, url, header, "item-4020102654.html")

	fmt.Printf("rsp=%v", rsp.Status())
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("------------add item to car---------------------\n")
	//request = client.R()
	//cookieStr := CopyCookies(rsp)
	//request.SetHeader("cookie", cookieStr)
	//AddCarRequestParam(rsp, request)
	//header = CommonPostHeader()
	//rsp2, _ := PostRequest(request, url, header, nil, "item-add-car.html")
	//
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("--------------look Cart----------------\n")
	//url2 := "https://www.anadf.com/cn/Cart.aspx"
	//request = client.R()
	//cookieStr = CopyCookies(rsp2)
	//request.SetHeader("cookie", cookieStr)
	//header = CommonGetHeader()
	//GetRequest(request, url2, header, "Cart.html")
	//
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("-----get MemberLogin------\n")
	////url = "https://www.anadf.com/cn/MemberLogin.aspx"
	//url = "https://www.anadf.com/cn/MemberLogin.aspx?ReturnUrl=cart"
	//request = client.R()
	//request.SetHeader("cookie", cookieStr)
	//header = CommonGetHeader()
	//rsp, _ = GetRequest(request, url, header, "MemberLogin.html")
	//
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("------------login------------------\n")
	//request = client.R()
	//cookieStr = CopyCookies(rsp)
	//request.SetHeader("cookie", cookieStr)
	//LoginRequestParam(rsp, request)
	//header = CommonPostHeader()
	//rsp, _ = PostRequest(request, url, header, nil, "MemberLogin-result.html")
	//
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("-----get ReserveEntry info------\n")
	//url = "https://www.anadf.com/cn/ReserveEntry.aspx"
	//request = client.R()
	//cookieStr = CopyCookies(rsp)
	//request.SetHeader("cookie", cookieStr)
	//header = CommonGetHeader()
	//rsp, _ = GetRequest(request, url, header, "ReserveEntry.html")
	//
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("-----post ReserveEntry info-----\n")
	//url = "https://www.anadf.com/cn/ReserveEntry.aspx"
	//request = client.R()
	//cookieStr = CopyCookies(rsp)
	//request.SetHeader("cookie", cookieStr)
	//PostCustomerInfoRequestParam(rsp, request)
	//header = CommonPostHeader()
	//rsp, _ = PostRequest(request, url, header, nil, "ReserveEntry-result.html")
	//
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("-----get ReserveEntryConfirm info-----\n")
	//url = "https://www.anadf.com/cn/ReserveEntryConfirm.aspx"
	//request = client.R()
	//cookieStr = CopyCookies(rsp)
	//request.SetHeader("cookie", cookieStr)
	//header = CommonGetHeader()
	//GetRequest(request, url, header, "ReserveEntryConfirm.html")
	//
	//fmt.Printf("\n")
	//fmt.Printf("\n")
	//fmt.Printf("-----post ReserveEntryConfirm-----\n")
	//url = "https://www.anadf.com/cn/ReserveEntryConfirm.aspx"
	//request = client.R()
	//cookieStr = CopyCookies(rsp)
	//request.SetHeader("cookie", cookieStr)
	//PostReserveEntryConfirmRequestParam(rsp, request)
	//header = CommonPostHeader()
	//PostRequest(request, url, header, nil, "ReserveEntryConfirm.html")
	//
	//endTime := time.Now()
	//fmt.Printf("begin time:%v\n", endTime.String())

}

func AddCarRequestParam(response *resty.Response, request *resty.Request) {
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

	form := map[string]string{
		"__EVENTTARGET":        __EVENTTARGET,
		"__EVENTARGUMENT":      __EVENTARGUMENT,
		"__LASTFOCUS":          __LASTFOCUS,
		"__VIEWSTATEGENERATOR": __VIEWSTATEGENERATOR,
		"__VIEWSTATE":          __VIEWSTATE,
		"__EVENTVALIDATION":    __EVENTVALIDATION,

		"ctl00$ddlAirport":  ddlAirport,
		"ctl00$ddlLanguage": ddlLanguage,
		"ctl00$txtKeyword":  txtKeyword,

		"NUM":     "1",
		"airport": "01",
		"ctl00$ContentPlaceHolder1$ucModalSelectAirport$btnConfirm": "OK",
	}

	request.SetFormData(form)
}

func LoginRequestParam(response *resty.Response, request *resty.Request) {
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

		"ctl00$ddlAirport":                  "01",
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

func PostCustomerInfoRequestParam(response *resty.Response, request *resty.Request) {
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

	txtKeyword := dom.Find("div input[name='ctl00$txtKeyword']").Eq(0).Text()
	fmt.Printf("ctl00$txtKeyword=%v\n", txtKeyword)

	form := map[string]string{
		"__EVENTTARGET":        __EVENTTARGET,
		"__EVENTARGUMENT":      __EVENTARGUMENT,
		"__LASTFOCUS":          __LASTFOCUS,
		"__VIEWSTATEGENERATOR": __VIEWSTATEGENERATOR,
		"__VIEWSTATE":          __VIEWSTATE,
		"__EVENTVALIDATION":    __EVENTVALIDATION,

		"ctl00$txtKeyword": txtKeyword,

		"departureDate": "20210820",
		"ctl00$ContentPlaceHolder1$ddlStrDateTime": "06",
		"flightNumber": "NH001",
		"ctl00$ContentPlaceHolder1$txtVisitorName": "",
		"ctl00$ContentPlaceHolder1$btnConfirm":     "??????????????????",

		// "ctl00$ddlAirport":  "01",
		// "ctl00$ddlLanguage": ddlLanguage,
		// "ctl00$txtKeyword":                  txtKeyword,
		// "ctl00$ContentPlaceHolder1$txtMail": "getway@moran.cn", //   "sdsdw@126.com"
		// "ctl00$ContentPlaceHolder1$TxtPASS": "moranjiuye1",     // "123123ab"
		// "ctl00$ContentPlaceHolder1$txtMail":  "sdsdw@126.com", //"getway@moran.cn",    //"getway@moran.cn"
		// "ctl00$ContentPlaceHolder1$TxtPASS":  "123123ab",      //"moranjiuye1",
		// "ctl00$ContentPlaceHolder1$btnLogin": btnLogin,
	}

	request.SetFormData(form)
}

func PostReserveEntryConfirmRequestParam(response *resty.Response, request *resty.Request) {
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

	// __LASTFOCUS, _ := dom.Find("input#__LASTFOCUS").Eq(0).Attr("value")
	// fmt.Printf("__LASTFOCUS=%v\n", __LASTFOCUS)

	__VIEWSTATE, _ := dom.Find("input#__VIEWSTATE").Eq(0).Attr("value")
	fmt.Printf("__VIEWSTATE=%v\n", __VIEWSTATE)

	__VIEWSTATEGENERATOR, _ := dom.Find("div input#__VIEWSTATEGENERATOR").Eq(0).Attr("value")
	fmt.Printf("__VIEWSTATEGENERATOR=%v\n", __VIEWSTATEGENERATOR)

	__EVENTVALIDATION, _ := dom.Find("div input#__EVENTVALIDATION").Eq(0).Attr("value")
	fmt.Printf("__EVENTVALIDATION=%v\n", __EVENTVALIDATION)

	txtKeyword := dom.Find("div input[name='ctl00$txtKeyword']").Eq(0).Text()
	fmt.Printf("ctl00$txtKeyword=%v\n", txtKeyword)

	form := map[string]string{
		"__EVENTTARGET":   __EVENTTARGET,
		"__EVENTARGUMENT": __EVENTARGUMENT,
		// "__LASTFOCUS":          __LASTFOCUS,
		"__VIEWSTATEGENERATOR": __VIEWSTATEGENERATOR,
		"__VIEWSTATE":          __VIEWSTATE,
		"__EVENTVALIDATION":    __EVENTVALIDATION,

		"ctl00$txtKeyword":                     txtKeyword,
		"ctl00$ContentPlaceHolder1$btnConfirm": "??????????????????",
	}

	request.SetFormData(form)
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
	fmt.Printf("-----post response header end-----\n")
	fmt.Printf("-----post response cookies:\n")
	for i, cookie := range resp.Cookies() {
		fmt.Printf("%v,%v\n", i, cookie.String())
	}
	fmt.Printf("-----post response cookies end\n")

	ioutil.WriteFile(saveFile, resp.Body(), 0600)
	return resp, nil
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
	fmt.Printf("-----get response header ene--------\n")
	fmt.Printf("-----get response cookies:\n")
	for i, cookie := range resp.Cookies() {
		fmt.Printf("%v,%v\n", i, cookie.String())
	}
	fmt.Printf("-----get response cookies end\n")

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

func CommonGetHeader() map[string]string {
	var header = map[string]string{

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
	return header
}

func CommonPostHeader() map[string]string {
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
	return header
}
