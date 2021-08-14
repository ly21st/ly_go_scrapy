package main

import (
	"crypto/tls"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"net/http"
	"yannscrapy/logging"
)



func main() {
	GetMyPage()
}

func GetMyPage() {
	url := "https://www.anadf.com/cn/MyPage.aspx"

	header := map[string]string {
		"accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		//"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN,zh;q=0.9",
		"cache-control": "max-age=0",
		"content-type": "application/x-www-form-urlencoded",
		"origin": "https://www.anadf.com",
		"referer": "https://www.anadf.com/cn/MemberLogin.aspx",
		"sec-ch-ua": `" Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`,
		"sec-ch-ua-mobile": "?0",
		"sec-fetch-dest": "document",
		"sec-fetch-mode": "navigate",
		"sec-fetch-site": "same-origin",
		"sec-fetch-user": "?1",
		"upgrade-insecure-requests": "1",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
		"cookie": `check_cookie=true; ASP.NET_SessionId=xkdt0aqt15phml43qchl3e1c; _ga=GA1.2.1652256127.1627650499; _rcmdjp_user_id=www.anadf.com-1346743333; check_cookie=true; visitorid=20210730220818358570; ADFWEB_CloseAlertAirportModal2=True; MasterLanguageType=2; ADFWEB_SessionId=; KeyWord_HEADER=; SCD=; Brand=; KeyWord=; Cate_Sub=; Cate_S=; Cate_M=; Cate_L=; PriceUpper=; PriceLower=; Recommended=; DFonly=; NewItem=; GroupCode=; ViewDate=; IsForce=; KeisaiFlg=; LimitedQty=; VariationCode=; ShowItemOfNoAirport=False; MasterMenuAirport2=04; QueryString=Search1Airport=01; AirPort=; Sort=0; token_2_0=UKjcx+mjju8HuXUpNoYZYTAVMbxTLm3mag/djJiVmy5ZkxbPrR0EMzNiSFMWovp4Juk5aUnjlNF3NEaprK2JZ1rZUZca+pHoeqANknicdIj2rz3HbgRAYXnRMeP4bYJhPckYAfyr1Hir8ju5jbJ46w==; _gid=GA1.2.1520344673.1628775541; _rcmdjp_history_view=4010102087_cn%2c4020102654_cn%2c4020101396_cn%2c4010102426_cn%2c4020102653_cn; JSESSIONID=D7A6C947C2C1F503D7A53FB00AA5892D; token_2_2=B1r4mm7varVMQwgjKJkjSyXVGYzUOcaZS0kYyi9Okv1PRf6Xb/oI8Xv3hw7NtrdncJcSKyRaCzKhI6F8g5mxlsnz10A6cKZGeb1poNjPEjJ0GBMadVNc4Dhu2hCY0d9La2YpCE8i/fR+XTnuKBT6dw==; _gat=1; .ADFWEBFORMSAUTH=C7D386541B8F7A152B790018C6B66C01F3E340058D676E618156FB6D85872CCB9D44940BD77B77EEA2E5A78D719043CCC352352F41E98A092067284170DE4510D74977755BCD1D4FE3E3D452DAAA699249305A71758F45453C0B5E1D11D838230C58DC87`,
	}

	//request := resty.R()
	//client := resty.New().SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true }).SetTimeout(60)
	client := resty.New().SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
	client.GetClient().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
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

	dstFileName := "mypage.html"
	ioutil.WriteFile(dstFileName, resp.Body(), 0600)

}
func SetHeader(request *resty.Request, m map[string]string) {
	for k, v := range m {
		request.SetHeader(k, v)
	}
}
