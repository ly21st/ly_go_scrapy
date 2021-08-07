package main

import (
	"yannscrapy/logging"
	"yannscrapy/utils"
)

func main() {

	//var speed, status = ProxyTest("http://124.205.155.151:9090")
	//var speed, status = ProxyTest("http://05.252.161.48:8080")

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
