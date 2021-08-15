package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

var cookies []*http.Cookie
var client = &http.Client{}

func init() {
	//跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Timeout: 30 * time.Second, Transport: tr}
}

func main() {
	// 用户客户端程序登录后，携带cookie，js-ajax跨域返回数据
	http.HandleFunc("/getData", getData)
	// 用户客户端程序登录后，携带cookie，重新定向URL页面
	http.HandleFunc("/login", doLogin)

	http.ListenAndServeTLS(":9966", "cert.pem", "key.pem", nil)
}

func doLogin(w http.ResponseWriter, r *http.Request) {
	if login() {
		http.SetCookie(w, cookies[0])
		http.Redirect(w, r, "http://127.0.0.1:997/index#/home", http.StatusFound)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	// 处理js-ajax跨域问题
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json

	if login() {
		resp2, err := client.Get("http://127.0.0.1:997/user/getTestData")
		defer resp2.Body.Close()
		if err != nil {
			err := fmt.Errorf("登录后发起Get请求时，client.Get错误:%v", err)
			fmt.Println(err)
			return
		}
		if resp2.StatusCode != 200 {
			err := fmt.Errorf("登录后发起Get请求时，response.StatusCode:%v", resp2.StatusCode)
			fmt.Println(err)
			return
		}

		body, err := ioutil.ReadAll(resp2.Body)

		w.Write(body)
	}

}

func login() bool {
	form := url.Values{}
	form.Set("username", "administrator")
	form.Set("password", "xxxxx")
	b := bytes.NewBufferString(form.Encode())

	req, err := http.NewRequest("POST", "http://127.0.0.1:997/login", b)
	if err != nil {
		err := fmt.Errorf("登录发起post请求时，http.NewRequest错误:%v", err)
		fmt.Println(err)
		return false
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		err := fmt.Errorf("登录发起post请求时，client.Do错误:%v", err)
		fmt.Println(err)
		return false
	}
	if resp.StatusCode != 200 {
		err := fmt.Errorf("登录发起post请求时，response.StatusCode:%v", resp.StatusCode)
		fmt.Println(err)
		return false
	}

	cookies = resp.Cookies()

	return true
}
