package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"yannscrapy/logging"
)


/**
检测代理ip
testUrl = "https://icanhazip.com"
testUrl = "http://icanhazip.com"
 */
func CheckProxyIp(proxyAddr string,
				httpsTestUrl string,
				httpTestUrl string,
				maxIdleConnsPerHost int,
				responseHeaderTimeout time.Duration,
				clientTimeout time.Duration) (Speed int, Status int) {
	if maxIdleConnsPerHost == 0 {
		maxIdleConnsPerHost = 5
	}

	if responseHeaderTimeout == 0 {
		responseHeaderTimeout = 5
	}

	if clientTimeout == 0 {
		clientTimeout = 5
	}

	//检测代理iP访问地址
	var testUrl string
	//判断传来的代理IP是否是https
	if strings.Contains(proxyAddr, "https") {
		testUrl = httpsTestUrl
	} else {
		testUrl = httpTestUrl
	}
	// 解析代理地址
	proxy, err := url.Parse(proxyAddr)
	//设置网络传输
	netTransport := &http.Transport{
		Proxy:                 http.ProxyURL(proxy),
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
		ResponseHeaderTimeout: time.Second * time.Duration(responseHeaderTimeout),
		TLSClientConfig: &tls.Config {
			InsecureSkipVerify: true,
		},
	}

	// 跳过https证书方法2
	//netTransport := &http.Transport{
	//	Proxy:                 http.ProxyURL(proxy),
	//	MaxIdleConnsPerHost:   maxIdleConnsPerHost,
	//	ResponseHeaderTimeout: time.Second * time.Duration(responseHeaderTimeout),
	//	TLSClientConfig: &tls.Config {
	//		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	//			return nil
	//		},
	//	},
	//}
	// 创建连接客户端
	httpClient := &http.Client{
		Timeout:   time.Second * clientTimeout,
		Transport: netTransport,
	}
	begin := time.Now() //判断代理访问时间
	// 使用代理IP访问测试地址
	res, err := httpClient.Get(testUrl)

	if err != nil {
		logging.Error(err)
		return
	}
	defer res.Body.Close()
	speed := int(time.Now().Sub(begin).Nanoseconds() / 1000 / 1000) //ms
	//判断是否成功访问，如果成功访问StatusCode应该为200
	if res.StatusCode != http.StatusOK {
		logging.Error(err)
		return
	}
	return speed, res.StatusCode
}


/**
默认方式检查代理ip
 */
func DefaultCheckProxyIp(proxyAddr string) (Speed int, Status int) {
	return CheckProxyIp(proxyAddr,
		"https://icanhazip.com",
		"http://icanhazip.com",
		0,
		0,
		0)
}

/**
获取代理ip
 */
func GetProxyIps(url string) ([]string, error) {
	httpRequest := new(HttpRequest)
	httpRequest.Timeout = 30
	response, err := httpRequest.Request(http.MethodGet, url)
	if err != nil {
		logging.Fatalf("request failed:%v\n", err)
		os.Exit(1)
	}

	ipList := make([]string, 0)
	bodyMap := make(map[string]interface{})
	err = json.Unmarshal(response.Body(), &bodyMap)
	if err != nil {
		logging.Error(err)
		return ipList, err
	}

	//logging.Infof("type(bodyMap)=%T", bodyMap)
	//logging.Infof("bodyMap=%v", bodyMap)

	code, err := ReadJsonObject(bodyMap, "code")
	if err != nil {
		logging.Error(err)
		return ipList, err
	}
	if code != float64(0) {
		logging.Errorf("response body code is not 0")
		return ipList, errors.New("response body code is not 0")
	}

	dataData, err := ReadJsonObject(bodyMap, "data.data")
	if err != nil {
		logging.Errorf("response body no data field")
		return ipList, errors.New("response body no data field")
	}

	dataDataList, ok := dataData.([]interface{})
	if !ok {
		msg := "dataData.([]interface{} error"
		logging.Errorf(msg)
		return ipList, errors.New(msg)
	}
	for _, dataVal := range dataDataList {
		proto, err := ReadJsonObject(dataVal, "protocol")
		if err != nil {
			logging.Errorf("utils.ReadJsonObject(dataVal, \"protocol\") error")
			break
		}

		ip, err := ReadJsonObject(dataVal, "ip")
		if err != nil {
			logging.Errorf("utils.ReadJsonObject(dataVal, \"ip\") error")
			break
		}

		port, err := ReadJsonObject(dataVal, "port")
		if err != nil {
			logging.Errorf("utils.ReadJsonObject(dataVal, \"port\") error")
			break
		}

		ipVal, _ := ip.(string)
		portVal, _ := port.(string)
		if proto == "http" {
			proxyHttpAddr := "http://" + ipVal + ":" + portVal
			ipList = append(ipList, proxyHttpAddr)
		} else if proto == "https" {
			proxyHttpAddr := "https://" + ipVal + ":" + portVal
			ipList = append(ipList, proxyHttpAddr)
		}
	}

	return ipList, nil
}


func GetAllHealthProxyIps(url string) ([]string, error){
	var result = make([]string, 0)
	var err error

	ipList, err := GetProxyIps(url)
	if err != nil {
		logging.Error(err)
		return result, err
	}

	for _, proxyAddr := range ipList {
		var speed, status = DefaultCheckProxyIp(proxyAddr)
		if status == 200 {
			logging.Infof("%s %d ms %d", proxyAddr, speed, status)
			result = append(result, proxyAddr)
		} else {
			logging.Infof("代理%s不可用", proxyAddr)
		}
	}

	return result, nil
}

func GetDefaultAllHealthProxyIps() ([]string, error){
	url := "https://ip.jiangxianli.com/api/proxy_ips"
	return GetAllHealthProxyIps(url)
}


func GetHealthProxyIp(url string) (string, error) {
	ipList, err := GetAllHealthProxyIps(url)
	if err != nil {
		logging.Error(err)
		return "", err
	}

	return ipList[rand.Intn(len(ipList))], nil
}


func GetDefaultHealthProxyIp() (string, error) {
	url := "https://ip.jiangxianli.com/api/proxy_ips"
	return GetHealthProxyIp(url)
}