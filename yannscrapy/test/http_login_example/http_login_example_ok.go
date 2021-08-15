package main

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	//"yannscrapy/logging"
)

//var header1 = map[string]string{
//
//	"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
//	"accept-language": "zh-CN,zh;q=0.9",
//	"cache-control": "max-age=0",
//
//	"sec-ch-ua":                 `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`,
//	"sec-ch-ua-mobile":          "?0",
//	"sec-fetch-dest":            "document",
//	"sec-fetch-mode":            "navigate",
//	"sec-fetch-site":            "same-origin",
//	"sec-fetch-user":            "?1",
//	"upgrade-insecure-requests": "1",
//	"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
//}

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
	//"cookie": `check_cookie=true; ASP.NET_SessionId=xkdt0aqt15phml43qchl3e1c; _ga=GA1.2.1652256127.1627650499; _rcmdjp_user_id=www.anadf.com-1346743333; check_cookie=true; visitorid=20210730220818358570; ADFWEB_CloseAlertAirportModal2=True; MasterLanguageType=2; ADFWEB_SessionId=; KeyWord_HEADER=; SCD=; Brand=; Cate_Sub=; Cate_S=; Cate_M=; Cate_L=; PriceUpper=; PriceLower=; Recommended=; DFonly=; NewItem=; GroupCode=; ViewDate=; IsForce=; KeisaiFlg=; LimitedQty=; VariationCode=; ShowItemOfNoAirport=False; QueryString=Search1Airport=01; AirPort=; token_2_0=UKjcx+mjju8HuXUpNoYZYTAVMbxTLm3mag/djJiVmy5ZkxbPrR0EMzNiSFMWovp4Juk5aUnjlNF3NEaprK2JZ1rZUZca+pHoeqANknicdIj2rz3HbgRAYXnRMeP4bYJhPckYAfyr1Hir8ju5jbJ46w==; _gid=GA1.2.1520344673.1628775541; JSESSIONID=9C5FEEB145BC14EF368A9BF0A6C130B6; MasterMenuAirport2=01; _rcmdjp_history_view=4020102654_cn,4010102087_cn,4020101396_cn,4010102426_cn,4020102653_cn; KeyWord=; Sort=0; token_2_2=KzG32vLVtekhYpA43oakQKTD/rKKhPsyAPcFY7gksPA2waY4iMteMhKwLr/c8KN9/xmToUxJMvgpnHFQWQedmIGkqv8b/z/BHZKFadgW99XjpABXvQ62kXOJeS/tS37sDbVc4PWdg5Lthbg4Nj+pbw==`,
	//"cookie": `check_cookie=true; _gid=GA1.2.1248265674.1628912072; _ga=GA1.2.137268062.1628912072; ADFWEB_SessionId=; KeyWord_HEADER=; SCD=; AirPort=; Brand=; KeyWord=; Cate_Sub=; Cate_S=; Cate_M=; Cate_L=; PriceUpper=; PriceLower=; Recommended=; DFonly=; NewItem=; GroupCode=; ViewDate=; IsForce=; KeisaiFlg=; LimitedQty=; VariationCode=; ShowItemOfNoAirport=False; Sort=0; ASP.NET_SessionId=wptxx2b4gt1vy2bwx5qdewpm; JSESSIONID=2D8AC933B543DDBC74596D534B90297C; visitorid=20210815163917547860; _rcmdjp_user_id=www.anadf.com-1199729120; ADFWEB_CloseAlertAirportModal2=True; MasterMenuAirport2=01; _gat=1; token_2_2=tuoALpV7XZ85IHDiwCrYPDIRP/335sX0b8m6G20wD6tzsorkgwnDuy7MGi97F/Sc1bE1GiYor7zF6DYsVEY6IXvHuK+RWIEk00zYLJ6a4Vw1RkcqV0Oim1JSDcovJ7qIVPEK07y/GsjsCoDhd5X3OA==`,

	//"cookie": `check_cookie=true; ASP.NET_SessionId=xkdt0aqt15phml43qchl3e1c; _ga=GA1.2.1652256127.1627650499; _rcmdjp_user_id=www.anadf.com-1346743333; check_cookie=true; visitorid=20210730220818358570; ADFWEB_CloseAlertAirportModal2=True; MasterLanguageType=2; ADFWEB_SessionId=; KeyWord_HEADER=; SCD=; Brand=; Cate_Sub=; Cate_S=; Cate_M=; Cate_L=; PriceUpper=; PriceLower=; Recommended=; DFonly=; NewItem=; GroupCode=; ViewDate=; IsForce=; KeisaiFlg=; LimitedQty=; VariationCode=; ShowItemOfNoAirport=False; QueryString=Search1Airport=01; AirPort=; token_2_0=UKjcx+mjju8HuXUpNoYZYTAVMbxTLm3mag/djJiVmy5ZkxbPrR0EMzNiSFMWovp4Juk5aUnjlNF3NEaprK2JZ1rZUZca+pHoeqANknicdIj2rz3HbgRAYXnRMeP4bYJhPckYAfyr1Hir8ju5jbJ46w==; _gid=GA1.2.1520344673.1628775541; KeyWord=; Sort=0; _rcmdjp_history_view=4020101396_cn,4020102654_cn,4010102087_cn,4010102426_cn,4020102653_cn; token_2_2=31Oehg7MTbsjmg43nXpDkZtBq2g1zk5kpDcP65t7R/eOdyOOHhTKdxWEuObbWiZLA262EQOCzdOcgVp23TJFBJHopC/L0kokCQPV2aXGeJ9y9SP/aXxPd1tcouMgB3fQoNI5X3+4zeD9zJjjaQTlAA==`,

   "cookie": `check_cookie=true; ASP.NET_SessionId=xkdt0aqt15phml43qchl3e1c; _ga=GA1.2.1652256127.1627650499; _rcmdjp_user_id=www.anadf.com-1346743333; check_cookie=true; visitorid=20210730220818358570; ADFWEB_CloseAlertAirportModal2=True; MasterLanguageType=2; ADFWEB_SessionId=; KeyWord_HEADER=; SCD=; Brand=; Cate_Sub=; Cate_S=; Cate_M=; Cate_L=; PriceUpper=; PriceLower=; Recommended=; DFonly=; NewItem=; GroupCode=; ViewDate=; IsForce=; KeisaiFlg=; LimitedQty=; VariationCode=; ShowItemOfNoAirport=False; QueryString=Search1Airport=01; AirPort=; token_2_0=UKjcx+mjju8HuXUpNoYZYTAVMbxTLm3mag/djJiVmy5ZkxbPrR0EMzNiSFMWovp4Juk5aUnjlNF3NEaprK2JZ1rZUZca+pHoeqANknicdIj2rz3HbgRAYXnRMeP4bYJhPckYAfyr1Hir8ju5jbJ46w==; _gid=GA1.2.1520344673.1628775541; KeyWord=; Sort=0; _rcmdjp_history_view=4020101396_cn,4020102654_cn,4010102087_cn,4010102426_cn,4020102653_cn; token_2_2=iLwa0N2GD0luFxc6X2nA6bCm36ZFTyepBa2XRTsboP8+VwwApo8tGGqLWHbqvU3mNq7ARStZFudJpNW+zVcgzPrCD0vhv0StEo8tvCW05X2qcA30EQb/WgWiO2T3wRXWQiVm8QRpi8ml4Xo36P+OVg==`,
}

func main() {
	client, err := CreateClient()
	if err != nil {
		fmt.Println(err)
	}
	//request := client.R()
	//url := "https://www.anadf.com/cn/MemberLogin.aspx"
	//rsp1, _ := GetRequest(request, url, header1, "MemberLogin.html")
	//
	//fmt.Printf("---------------------------------------\n")
	//fmt.Printf("---------------------------------------\n")
	//
	url := "https://www.anadf.com/cn/MemberLogin.aspx"
	//
	//fmt.Printf(" get response cookies:%v\n",rsp1.Cookies())

	//client.SetCookies(rsp1.Cookies())
	request := client.R()

	//copyRequestParam01(rsp1, request)
	copyRequestParam02(request)
	//request.SetHeader("")
	//request.SetHeader("cookie", `check_cookie=true; _gid=GA1.2.1248265674.1628912072; _ga=GA1.2.137268062.1628912072; ADFWEB_SessionId=; KeyWord_HEADER=; SCD=; AirPort=; Brand=; KeyWord=; Cate_Sub=; Cate_S=; Cate_M=; Cate_L=; PriceUpper=; PriceLower=; Recommended=; DFonly=; NewItem=; GroupCode=; ViewDate=; IsForce=; KeisaiFlg=; LimitedQty=; VariationCode=; ShowItemOfNoAirport=False; Sort=0; ASP.NET_SessionId=wptxx2b4gt1vy2bwx5qdewpm; JSESSIONID=2D8AC933B543DDBC74596D534B90297C; visitorid=20210815163917547860; _rcmdjp_user_id=www.anadf.com-1199729120; ADFWEB_CloseAlertAirportModal2=True; MasterMenuAirport2=01; _gat=1; token_2_2=tuoALpV7XZ85IHDiwCrYPDIRP/335sX0b8m6G20wD6tzsorkgwnDuy7MGi97F/Sc1bE1GiYor7zF6DYsVEY6IXvHuK+RWIEk00zYLJ6a4Vw1RkcqV0Oim1JSDcovJ7qIVPEK07y/GsjsCoDhd5X3OA==`)
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
		"__EVENTTARGET":   __EVENTTARGET,
		"__EVENTARGUMENT": __EVENTARGUMENT,
		"__LASTFOCUS":     __LASTFOCUS,
		"__VIEWSTATE":     __VIEWSTATE,

		"__VIEWSTATEGENERATOR": __VIEWSTATEGENERATOR,
		"__EVENTVALIDATION":    __EVENTVALIDATION,

		"ctl00$ddlAirport":                   "01",
		"ctl00$ddlLanguage":                  ddlLanguage,
		"ctl00$txtKeyword":                   txtKeyword,
		"ctl00$ContentPlaceHolder1$txtMail":  "sdsdw@126.com", //"getway@moran.cn",    //"getway@moran.cn"
		"ctl00$ContentPlaceHolder1$TxtPASS":  "123123ab",      //"moranjiuye1",
		"ctl00$ContentPlaceHolder1$btnLogin": btnLogin,
	}

	request.SetFormData(form)
}

func copyRequestParam01(response *resty.Response, request *resty.Request) {
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
		"__EVENTTARGET":        "",
		"__EVENTARGUMENT":      "",
		"__LASTFOCUS":          "",
		"__VIEWSTATEGENERATOR": "AA4F3242",
		//"__VIEWSTATE": __VIEWSTATE,
		//"__EVENTVALIDATION": __EVENTVALIDATION,

		//"__VIEWSTATE":       "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",
		//"__EVENTVALIDATION": "lCava3y/+MBfYZjO4PWwitIx9jY8hnkmm0jsUV/h13gscWJMqMMlNcgiTdd59NNtBOFg3G7HbMio57N7gtw3p0sAxECr6VyagnyUMGK35XexWMZsnwpYNyoG6aroYfZrLDma/kA2H0cjD7vr2dsZKb5Z/cz6La0pSsKOPMiQ6ALdR/o/kHR/PwB7UC59tKDrdMlJvnkOV7ftD9TcDrfSOAz9JtJ0umi8RjI+RbPYRcRlYxKOATihhBO6YlWcqDeWgotmn/TPhFOKjY+WN47b532gt4VfjxOpfi1GWIdriNpoiN6tkFEkBmYhYwtJmtDPnVHHaUDa4doWvY93pPy90cdF4UWod6jR6QuHg/jJvh7mMnSW8Y6IZdFB2KnJzn0J0Il+vECgLbWfKIWLHb1S4E/nPo88n4hNnR2Dt1nSMYKWTpFcWVG3vzCvN0TUynMviwfFHmt6vU6ItDPwO8tmD7C6svDBbsfBGpX16AoGJGlzIPqnDOjFzU8EbbqPIrxoIlshTmjnLf9aWdiqnVaROVnmO/4=",


		//"__VIEWSTATE": "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",
		//"__EVENTVALIDATION": "zvM0zDZrHtuUdxw9KpgOEDXsJEzOTGU/TmGcSxwYzOQ7RvIrpUQY497t3hh7k3p63BIl25i93wzAr/wHQXg+O8gExnnEyJGrSWsVm9CB2DPBSTMMgKkWND9HRkkWZp1gS0Lq7PK5B5rBkRy/N1OFSDeggBSBNYqV9tO0lYsh8gbIxELHjeCRWFPKJYMGZi+rOSB3tbW4WSITPuPXJd/SZNHf7SOEmjZJaYhykNTApVLDYWkIDniKx+xPOuluyGrDb4olbRwEbvRWX/YxzX++BFcbd1aVOEoEbnbFm0hsvm8qzIedyMf1GCOl8LmpM2FVP9J9R2jrf1J/wRaywN5QCKORmkyypyDV0S5AuonGTOsg2kK72PgCo7FLmYiu/JiRBLmefQvZ7VDbXkpDihIpvPJ5Tm1XK1c/13gFLTmwSqDTI5x/7ioNcaCQbpteJPDAgidXw5TNNI/asFRlB397ZtJ3jP4Ek5V3iVm377Rpj7x3UJbSDg5weCpvq7dr6lQKEdUYt+FdXt2t6nBfzEnS+xSO4Q0=",

		"ctl00$ddlAirport":  "01",
		"ctl00$ddlLanguage": "2",
		"ctl00$txtKeyword":  "",
		//"ctl00$ContentPlaceHolder1$txtMail":  "getway@moran.cn",     //   "sdsdw@126.com"
		//"ctl00$ContentPlaceHolder1$TxtPASS":  "moranjiuye1",          // "123123ab"
		"ctl00$ContentPlaceHolder1$txtMail":  "sdsdw@126.com", //"getway@moran.cn",    //"getway@moran.cn"
		"ctl00$ContentPlaceHolder1$TxtPASS":  "123123ab",      //"moranjiuye1",
		"ctl00$ContentPlaceHolder1$btnLogin": "登录",
	}

	request.SetFormData(form)
}

func copyRequestParam02(request *resty.Request) {

	form := map[string]string{
		"__EVENTTARGET":        "",
		"__EVENTARGUMENT":      "",
		"__LASTFOCUS":          "",
		"__VIEWSTATEGENERATOR": "AA4F3242",
		//"__VIEWSTATE": __VIEWSTATE,
		//"__EVENTVALIDATION": __EVENTVALIDATION,

		//"__VIEWSTATE":       "lUmaArU1esvoy/A+B7KGS/7S8u9moK4M8YcNFFW+RIW5Ei8F9xp7qK5Yn4+W1cyA5t3LDY7IM+nF4ajRZmhOYQ8/DSgmJ7OMr4ol5aSEOjVxYLjkWz1SLQNlwl2MD7HFCuzCFvRc6T/V5+z/q8nnPzYRriPMAEDubRqHzZCkicBKkhR9kYTCV3YRzDqiy6kArzH+rcmUaDOwLZI0YqmgWQmpg6PPNfwsWySTSGhy2KRepqj9iao+mGGEkifMqDH1BoT+u10u0UCeJDm5m3hMKdJATKPqvs7lzLgwWQsXUVg1qhBGvu+S3RNCuYOsk6D2M+m0encAAoFDCzD7lUIyiym6ERyG+QWI2A+FZe00TGZPcKL/lY7wpFJBJrnfafI2Y9Lxs5RJNt5bmIb/GUuZZGfYAwIKx9ea3DkvfaZAVgV7gmiA2IA5Ifb9BkVXFfwCTGhPF65eMYmd1n4VXIeW7yg1g45rdw7gScO2SJ6objlOxrppDQwKww9v/2C6fGPIYnqjIs1Zp3vBgBiZc5LjAweMGwq5TsV1zf0lwmrpjz47YMV/Sg1fDoXkjfrx3CyweY10Xdrx1jppLnF2XnKZ9DFzF3gpRx6ZWvSMc3JVZ4+vvHQhIa/AbtJyknU/ymXzmWU7uQwCY3sabvjDx2ev4BtbB1cPHpwLmaAcSgqWlYKo4R7Y5rbV1xrU4000+JVgalFQm3UZ8vIAJJssbeFlwyW+yJn9Hko7Z5q8hWwjBkGmZ/58heEGev2U7/a8RS7RZSYnZ0VDX3VmiIPumFqsEc4fc+JxGdnah/G90Lz6+09ZRKh486Jf0WZrEX063zoOBZ62jrr2bg+A8SDZObEShKtxdNSt1rdB2Eb0fRFf8x0Manq6FOQJPLkqX7g4sOpzjNRETgWOjeA8t9x5StgCwUBfw2RPbyhQgopZEqCcAdFd90rjD/ANCpV+BAHhcsRDpkjWoV3uoQ/hklvNsJadJnoHRu0abPlsSZkjW8kEz8d9VxCWTZrCm8QAiMTPCacpkr7rt7ZqTekql438hFtNvAd6B13G/CssVPOdkgrJyv+FWNO++YJwvZNLQ8DkZ+PoLo9WJES9dk9DwH5iHto+wPxVv08i5bUdtPMuifrfj/EPvq1C6iJ0OqE63w3iFPwVFLYmt1CSRzIpoS4l74MDipFvDPdXqdJApMYf+MT91ape+UTJjrPPJYip27lnOAGCrGYCCA6hDIQTxDZ0eiD4QkRtBmE9qMVh1EMCAIW/N3/85GORFvrA1t/2nRW0sOvYUVBs/o5amupERHDnkyY+ErDmptlUPe+Sl9uGlJFCQPXE2D6yh468BbYo2nYK1NT2ndhY68r/vzedFEm/RAXFyv4PmryMNvM4Wh7W3ly3spEF4USJ2KMqcAPx5cOk/caJdS6gPxdCE4oInrQsbClxOtTQGp1FHRixO7wXGG233vQV2MvXIrAEHn66q7DEzA8JizuU/4+2SFYCGWHuIdt9zyz+PWXEfLLPI4pUIuiZtjx9EnpdDUsddOmr5dJ4H96trMtkVDage96d2VeNEatvXbsJ2HKsMMK/i13uj66+OOEzgwZ1od9aLacp2KHb3tCL9JncKa5VAg+GY+CJM/FK+RAYHVOAiqQ7gq+IoTcsVQvEIXyHHQ409UjUq/RsF4xe+SVnpbgJnhhe3Klc8B1PqChJFUMU3Qc2f0PVRjOeiYjwLzoOnUKN1jbT1wxAw7lyh2eWmY7WOiZXSBVMuO/dvje5grrNhJ9baLiZbs4xrcvM00szjF4QndoK5rvRb7M97YWOyxwSR8v5mhhgKS5lqDdK1evwSqDReTH0bwTRexRb9KyXxxWdim6P1QH444kp6dVmvtCXIuDcmbWyqadA3v2l7lHE51ShfVLNIleRTictJ6z1tqx48v2XGOwuAHr2HWVzMCb0qqlg6BSWgZH5bmWzEEgJjxeQqLHlc0MIJHrOXziaLHwIblZkq2IdHd1YtWD0ObKYGW15bdDxv18nKmpaWNEnxQrnAEJb6mQowVGE3wstN361x9FXS905TtAEbYWQUld1FWG9X8lOTvatDX5Xu60AWtxl+QESj0V7o1CgEuhEAFYowUuTH7srbNN59bvI/zF2qz5qSV+hPcydmNZmYlrB439xSyypWTsG6fRpENvKQqwwSvgXKxoxdyCSo1WwplKIqhwHhkUXKeVbVDTx43ZuVzD3eV9tdXefnms8jZYMd86iyEt6q+GkIRz4vyzXXoW+QxxqMFmacEqVJ/CTGq7a7wPLilYLwvajLZb6PyhOUDCJ3SboznEbyOVyrgXYd8GjoH3/YS7kS5WgO4rAEe4Z7n0paCiXJ/srxD9lhVBjBZbKYkodKH1fkdlWlbMgbw==",
		//"__EVENTVALIDATION": "lCava3y/+MBfYZjO4PWwitIx9jY8hnkmm0jsUV/h13gscWJMqMMlNcgiTdd59NNtBOFg3G7HbMio57N7gtw3p0sAxECr6VyagnyUMGK35XexWMZsnwpYNyoG6aroYfZrLDma/kA2H0cjD7vr2dsZKb5Z/cz6La0pSsKOPMiQ6ALdR/o/kHR/PwB7UC59tKDrdMlJvnkOV7ftD9TcDrfSOAz9JtJ0umi8RjI+RbPYRcRlYxKOATihhBO6YlWcqDeWgotmn/TPhFOKjY+WN47b532gt4VfjxOpfi1GWIdriNpoiN6tkFEkBmYhYwtJmtDPnVHHaUDa4doWvY93pPy90cdF4UWod6jR6QuHg/jJvh7mMnSW8Y6IZdFB2KnJzn0J0Il+vECgLbWfKIWLHb1S4E/nPo88n4hNnR2Dt1nSMYKWTpFcWVG3vzCvN0TUynMviwfFHmt6vU6ItDPwO8tmD7C6svDBbsfBGpX16AoGJGlzIPqnDOjFzU8EbbqPIrxoIlshTmjnLf9aWdiqnVaROVnmO/4=",


		//"__VIEWSTATE": "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",
		//"__EVENTVALIDATION": "zvM0zDZrHtuUdxw9KpgOEDXsJEzOTGU/TmGcSxwYzOQ7RvIrpUQY497t3hh7k3p63BIl25i93wzAr/wHQXg+O8gExnnEyJGrSWsVm9CB2DPBSTMMgKkWND9HRkkWZp1gS0Lq7PK5B5rBkRy/N1OFSDeggBSBNYqV9tO0lYsh8gbIxELHjeCRWFPKJYMGZi+rOSB3tbW4WSITPuPXJd/SZNHf7SOEmjZJaYhykNTApVLDYWkIDniKx+xPOuluyGrDb4olbRwEbvRWX/YxzX++BFcbd1aVOEoEbnbFm0hsvm8qzIedyMf1GCOl8LmpM2FVP9J9R2jrf1J/wRaywN5QCKORmkyypyDV0S5AuonGTOsg2kK72PgCo7FLmYiu/JiRBLmefQvZ7VDbXkpDihIpvPJ5Tm1XK1c/13gFLTmwSqDTI5x/7ioNcaCQbpteJPDAgidXw5TNNI/asFRlB397ZtJ3jP4Ek5V3iVm377Rpj7x3UJbSDg5weCpvq7dr6lQKEdUYt+FdXt2t6nBfzEnS+xSO4Q0=",

		"__VIEWSTATE": `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`,
		"__EVENTVALIDATION": `ePI7NUihtpEyI9KCSTg9ys3BXD5c66a5QU/CFLeqdK5g81/MTWKkMdF4z8fXBWHHhYn0mxhb1lNIck7XSZ5n+s1qsC573Qd+LpuvJQ9i+SQDGTcqMzT2wdlmZAuKofgpCiHnR9He+2OigpxqQL/Bj6sl5VohkM1GyBXpNZ5H6Ffc1YThvhEVmXUUNormn8uT4Fby+MFECLPspASydsA2GGYVuSgo6VhNkW2b6gvUWi9HLOLsY529fYLhT8qbqmjIWcN+GcBKYhDQaULMUBvUxptVaJMKUfAwlr37Q7YbZkd2fw+Iggos+L5ZJzl9yTkvaS5gQQDUc/okUTQ95HxC3DOQzpq8yoaQvoj3jHMJQdZlRBcNrTtbAYsaknrFw0nIbE5gGOhapLTqfBSk4OhGUEMjVg9Z9Dek0VkJahGHBW+P8taEMStGrxzt5sKRpSCSs1uiTNxKXXXsvOTImBk4FFt+1ouiktKoYg+9aBF4vV1OLejo2MPZ665y8LxKAsPQXRsVcUXhr75WlTiL7NttXoIxRVI=`,


		"ctl00$ddlAirport":  "00",
		//"ctl00$ddlAirport":  "01",
		"ctl00$ddlLanguage": "2",
		"ctl00$txtKeyword":  "",
		"ctl00$ContentPlaceHolder1$txtMail":  "getway@moran.cn",
		"ctl00$ContentPlaceHolder1$TxtPASS":  "moranjiuye1",
		//"ctl00$ContentPlaceHolder1$txtMail":  "sdsdw@126.com",
		//"ctl00$ContentPlaceHolder1$TxtPASS":  "123123ab",
		"ctl00$ContentPlaceHolder1$btnLogin": "登录",
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
