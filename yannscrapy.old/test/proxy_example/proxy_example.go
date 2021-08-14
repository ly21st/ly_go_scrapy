package main

import (
	"yannscrapy/logging"
	"yannscrapy/utils"
)

func main() {

	//test1()
	//test2()

	ip, err := utils.GetDefaultHealthProxyIp()
	if err != nil {
		logging.Fatal(err)
	}
	logging.Infof("ip=%s", ip)

}

func test1() {
	url := "https://ip.jiangxianli.com/api/proxy_ips"
	ipList, err := utils.GetProxyIps(url)
	if err != nil {
		logging.Error(err)
	}

	for _, proxyAddr := range ipList {
		var speed, status = utils.DefaultCheckProxyIp(proxyAddr)
		if status == 200 {
			logging.Infof("%s %d ms %d", proxyAddr, speed, status)
		} else {
			logging.Infof("代理%s不可用", proxyAddr)
		}
	}

}


func test2() {
	url := "https://ip.jiangxianli.com/api/proxy_ips"
	ipList, err := utils.GetAllHealthProxyIps(url)
	if err != nil {
		logging.Fatal(err)
	}
	for _, ip := range ipList {
		logging.Infof("ip=%s", ip)
	}
}