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

	// 搜索商品
	url := "https://www.anadf.com/cn/ItemDetail.aspx?S_CD="
	url = url + "4020102654"
	rsp, _, _ := SearchItem(client, url, "4020102654")

	fmt.Printf("rsp=%v", rsp.Status())

	//  添加商品
	cookieStr := CopyCookies(rsp)
	rsp, _, _ = AddItemToCar(client, url, rsp, cookieStr, "01")
	fmt.Printf("rsp=%v", rsp.Status())

	//
	// 查看购物车
	cookieStr = CopyCookies(rsp)
	url = "https://www.anadf.com/cn/Cart.aspx"
	LookCar(client, url, cookieStr)

	// 查看登录页面
	url = "https://www.anadf.com/cn/MemberLogin.aspx?ReturnUrl=cart"
	rsp, _, _ = GetMemberLogin(client, url, cookieStr)

	// 登录
	cookieStr = CopyCookies(rsp)
	url = "https://www.anadf.com/cn/MemberLogin.aspx?ReturnUrl=cart"
	rsp, _, _ = Login(client, url, rsp, cookieStr)

	// 查看登录后预约页
	cookieStr = CopyCookies(rsp)
	url = "https://www.anadf.com/cn/ReserveEntry.aspx"
	rsp, _, _ = GetReserveEntry(client, url, cookieStr)

	// 提交预约信息
	cookieStr = CopyCookies(rsp)
	url = "https://www.anadf.com/cn/ReserveEntry.aspx"
	rsp, _, _ = PostReserveEntry(client, url, rsp, cookieStr,
		"20210821", "06", "NH001", "")

	//  查看预约确认
	cookieStr = CopyCookies(rsp)
	url = "https://www.anadf.com/cn/ReserveEntryConfirm.aspx"
	rsp, _, _ = GetReserveEntryConfirm(client, url, cookieStr)


	// 提交预约确认
	//cookieStr = CopyCookies(rsp)
	url = "https://www.anadf.com/cn/ReserveEntryConfirm.aspx"
	PostReserveEntryConfirm(client, url, rsp, cookieStr)

	endTime := time.Now()
	fmt.Printf("end time:%v, cost:%vs\n", endTime.String(), time.Since(beginTime))

}

func SearchItem(client *resty.Client, url string, itemId string) (*resty.Response, int, error) {
	fmt.Printf("\n\n")
	fmt.Printf("-----search item-----")

	request := client.R()
	header := CommonGetHeader()
	filename := "item-" + itemId + ".html"
	rsp, err := GetRequest(request, url, header, filename)

	return rsp, rsp.StatusCode(), err
}

func AddItemToCar(client *resty.Client, url string, lastRsp *resty.Response, cookieStr string, airPort string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("------------add item to car---------------------\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	AddCarRequestParam(lastRsp, request, airPort)
	header := CommonPostHeader()
	rsp, err := PostRequest(request, url, header, nil, "item-add-car.html")
	return rsp, rsp.StatusCode(), err
}

func LookCar(client *resty.Client, url string, cookieStr string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("--------------look Cart----------------\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	header := CommonGetHeader()
	rsp, err := GetRequest(request, url, header, "Cart.html")
	return rsp, rsp.StatusCode(), err
}

func GetMemberLogin(client *resty.Client, url string, cookieStr string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("-----get MemberLogin------\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	header := CommonGetHeader()
	rsp, err := GetRequest(request, url, header, "MemberLogin.html")
	return rsp, rsp.StatusCode(), err
}

func Login(client *resty.Client, url string, lastRsp *resty.Response, cookieStr string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("------------login------------------\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	LoginRequestParam(lastRsp, request)
	header := CommonPostHeader()
	rsp, err := PostRequest(request, url, header, nil, "MemberLogin-result.html")
	return rsp, rsp.StatusCode(), err
}

func GetReserveEntry(client *resty.Client, url string, cookieStr string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("-----get ReserveEntry info------\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	header := CommonGetHeader()
	rsp, err := GetRequest(request, url, header, "ReserveEntry.html")
	return rsp, rsp.StatusCode(), err
}

func PostReserveEntry(client *resty.Client, url string, lastRsp *resty.Response, cookieStr string,
	departureDate string,
	ddlStrDateTime string,
	flightNumber string,
	txtVisitorName string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("-----post ReserveEntry info-----\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	PostCustomerInfoRequestParam(lastRsp, request,
		departureDate,
		ddlStrDateTime,
		flightNumber,
		txtVisitorName)
	header := CommonPostHeader()
	rsp, err := PostRequest(request, url, header, nil, "ReserveEntry-result.html")
	return rsp, rsp.StatusCode(), err
}

func GetReserveEntryConfirm(client *resty.Client, url string, cookieStr string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("-----get ReserveEntryConfirm info-----\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	header := CommonGetHeader()
	rsp, err := GetRequest(request, url, header, "ReserveEntryConfirm.html")
	return rsp, rsp.StatusCode(), err
}

func PostReserveEntryConfirm(client *resty.Client, url string, lastRsp *resty.Response, cookieStr string) (*resty.Response, int, error) {
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("-----post ReserveEntryConfirm-----\n")
	request := client.R()
	request.SetHeader("cookie", cookieStr)
	PostReserveEntryConfirmRequestParam(lastRsp, request)
	header := CommonPostHeader()
	rsp, err := PostRequest(request, url, header, nil, "ReserveEntryConfirm.html")
	return rsp, rsp.StatusCode(), err
}

func AddCarRequestParam(response *resty.Response, request *resty.Request, airPort string) {
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
		"airport": airPort,
		"ctl00$ContentPlaceHolder1$ucModalSelectAirport$btnConfirm": "OK",

		//"ctl00$ScriptManager1": "ctl00$ContentPlaceHolder1$UpdatePanel_javascript|ctl00$ContentPlaceHolder1$BtnAddCart",
		//"ctl00$ddlAirport": airPort,
		//"ctl00$ddlLanguage": "2"
		//"ctl00$txtKeyword": txtKeyword,
		//"NUM": "1"
		//"__EVENTTARGET": "ctl00$ContentPlaceHolder1$BtnAddCart",
		//"__EVENTARGUMENT":"",
		//"__LASTFOCUS":"",
		//"__VIEWSTATE":"",
		//
		//"__VIEWSTATEGENERATOR": "E2F80F2D",
		//"__EVENTVALIDATION": "",
		//"__ASYNCPOST": "true",

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

func PostCustomerInfoRequestParam(response *resty.Response, request *resty.Request,
	departureDate string,
	ddlStrDateTime string,
	flightNumber string,
	txtVisitorName string) {
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

	//__LASTFOCUS, _ := dom.Find("input#__LASTFOCUS").Eq(0).Attr("value")
	//fmt.Printf("__LASTFOCUS=%v\n", __LASTFOCUS)

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
		//"__LASTFOCUS":          __LASTFOCUS,
		"__VIEWSTATEGENERATOR": __VIEWSTATEGENERATOR,
		"__VIEWSTATE":          __VIEWSTATE,
		"__EVENTVALIDATION":    __EVENTVALIDATION,

		"ctl00$txtKeyword": txtKeyword,
		"departureDate":    departureDate,
		"ctl00$ContentPlaceHolder1$ddlStrDateTime": ddlStrDateTime,
		"flightNumber": flightNumber,
		"ctl00$ContentPlaceHolder1$txtVisitorName": txtVisitorName,
		"ctl00$ContentPlaceHolder1$btnConfirm":     "确认输入内容",
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
		"ctl00$ContentPlaceHolder1$btnConfirm": "进行预约申请",
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

	//ip, _ := utils.GetDefaultHealthProxyIp()
	//fmt.Printf("ip=%v", ip)
	//proxy, err := url.Parse(ip)
	//ip := "124.121.2.160"
	////http.ProxyURL(proxy)
	//client.SetProxy(ip)

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
