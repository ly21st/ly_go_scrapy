package main

import (
	"crypto/tls"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
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
	"authority": "www.anadf.com",
	"method": "POST",
	"path": "/cn/MemberLogin.aspx",
	"scheme": "https",

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
	//"cookie": `check_cookie=true; ASP.NET_SessionId=xkdt0aqt15phml43qchl3e1c; _ga=GA1.2.1652256127.1627650499; _rcmdjp_user_id=www.anadf.com-1346743333; check_cookie=true; visitorid=20210730220818358570; ADFWEB_CloseAlertAirportModal2=True; MasterLanguageType=2; ADFWEB_SessionId=; KeyWord_HEADER=; SCD=; Brand=; Cate_Sub=; Cate_S=; Cate_M=; Cate_L=; PriceUpper=; PriceLower=; Recommended=; DFonly=; NewItem=; GroupCode=; ViewDate=; IsForce=; KeisaiFlg=; LimitedQty=; VariationCode=; ShowItemOfNoAirport=False; QueryString=Search1Airport=01; AirPort=; token_2_0=UKjcx+mjju8HuXUpNoYZYTAVMbxTLm3mag/djJiVmy5ZkxbPrR0EMzNiSFMWovp4Juk5aUnjlNF3NEaprK2JZ1rZUZca+pHoeqANknicdIj2rz3HbgRAYXnRMeP4bYJhPckYAfyr1Hir8ju5jbJ46w==; _gid=GA1.2.1520344673.1628775541; JSESSIONID=9C5FEEB145BC14EF368A9BF0A6C130B6; MasterMenuAirport2=01; _rcmdjp_history_view=4020102654_cn,4010102087_cn,4020101396_cn,4010102426_cn,4020102653_cn; KeyWord=; Sort=0; token_2_2=KzG32vLVtekhYpA43oakQKTD/rKKhPsyAPcFY7gksPA2waY4iMteMhKwLr/c8KN9/xmToUxJMvgpnHFQWQedmIGkqv8b/z/BHZKFadgW99XjpABXvQ62kXOJeS/tS37sDbVc4PWdg5Lthbg4Nj+pbw==`,

}


func main() {
	client, err := CreateClient()
	if err != nil {
		logging.Panic(err)
	}
	request := client.R()
	url := "https://www.anadf.com/cn/MemberLogin.aspx"
	rsp1, _ := GetRequest(request, url, header1, "MemberLogin.html")

	logging.Infof("---------------------------------------")
	logging.Infof("---------------------------------------")


	//client.SetCookies(rsp1.Cookies())
	request = client.R()
	copyRequestParam(rsp1, request)
	PostRequest(request, url, header, nil, "login_result.html")
}

func copyRequestParam(response *resty.Response, request *resty.Request) {
	//dom, err := goquery.NewDocumentFromReader(rsp1.RawBody())
	//logging.Infof("body=%v", string(rsp1.Body()))
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(response.Body())))
	if err != nil {
		logging.Panic(err)
	}

	__EVENTTARGET, _ := dom.Find("input#__EVENTTARGET").Eq(0).Attr("value")
	logging.Infof("__EVENTTARGET=%v", __EVENTTARGET)

	__EVENTARGUMENT,_ := dom.Find("input#__EVENTARGUMENT").Eq(0).Attr("value")
	logging.Infof("__EVENTARGUMENT=%v", __EVENTARGUMENT)

	__LASTFOCUS,_ := dom.Find("input#__LASTFOCUS").Eq(0).Attr("value")
	logging.Infof("__LASTFOCUS=%v", __LASTFOCUS)

	__VIEWSTATE,_ := dom.Find("input#__VIEWSTATE").Eq(0).Attr("value")
	logging.Infof("__VIEWSTATE=%v", __VIEWSTATE)

	__VIEWSTATEGENERATOR, _ := dom.Find("div input#__VIEWSTATEGENERATOR").Eq(0).Attr("value")
	logging.Infof("__VIEWSTATEGENERATOR=%v", __VIEWSTATEGENERATOR)

	__EVENTVALIDATION, _ := dom.Find("div input#__EVENTVALIDATION").Eq(0).Attr("value")
	logging.Infof("__EVENTVALIDATION=%v", __EVENTVALIDATION)


	ddlAirport, _ := dom.Find("div select[name='ctl00$ddlAirport'] option[selected=selected]").Eq(0).Attr("value")
	logging.Infof("ctl00$ddlAirport=%v", ddlAirport)

	ddlLanguage, _ := dom.Find("div select[name='ctl00$ddlLanguage'] option[selected=selected]").Eq(0).Attr("value")
	logging.Infof("ctl00$ddlLanguage=%v", ddlLanguage)

	txtKeyword := dom.Find("div input[name='ctl00$txtKeyword']").Eq(0).Text()
	logging.Infof("ctl00$txtKeyword=%v", txtKeyword)

	txtMail := dom.Find("div input[name='ctl00$ContentPlaceHolder1$txtMail']").Eq(0).Text()
	logging.Infof("ctl00$ContentPlaceHolder1$txtMail=%v", txtMail)

	txtPass := dom.Find("div input[name='ctl00$ContentPlaceHolder1$TxtPASS']").Eq(0).Text()
	logging.Infof("ctl00$ContentPlaceHolder1$TxtPASS=%v", txtPass)

	btnLogin, _ := dom.Find("div input[name='ctl00$ContentPlaceHolder1$btnLogin']").Eq(0).Attr("value")
	logging.Infof("ctl00$ContentPlaceHolder1$btnLogin=%v", btnLogin)

	form := map[string]string {
		"__EVENTTARGET":        __EVENTTARGET,
		"__EVENTARGUMENT":      __EVENTARGUMENT,
		"__LASTFOCUS":          __LASTFOCUS,
		"__VIEWSTATE":          __VIEWSTATE,

		"__VIEWSTATEGENERATOR": __VIEWSTATEGENERATOR,
		"__EVENTVALIDATION":    __EVENTVALIDATION,

		"ctl00$ddlAirport":                   "01",
		"ctl00$ddlLanguage":                  ddlLanguage,
		"ctl00$txtKeyword":                   txtKeyword,
		"ctl00$ContentPlaceHolder1$txtMail":  "getway@moran.cn",
		"ctl00$ContentPlaceHolder1$TxtPASS":  "moranjiuye1",
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

	resp, err := request.Post(url)
	if err != nil {
		logging.Panic(err)
	}

	logging.Infof("--------------------------------------")
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

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.GetClient().Jar = jar
	return client, nil
}


func GetRequest(request *resty.Request, url string, header map[string]string, saveFile string) (*resty.Response, error){
	if header != nil {
		SetHeader(request, header)
	}
	resp, err := request.Get(url)
	if err != nil {
		logging.Panic(err)
	}
	logging.Infof("--------------------------------------")
	logging.Infof("status=%v", resp.Status())
	logging.Infof("------header-----")
	for e, v := range resp.Header() {
		logging.Infof("%v:%v", e, v)
	}

	ioutil.WriteFile(saveFile, resp.Body(), 0600)
	return resp, nil
}
