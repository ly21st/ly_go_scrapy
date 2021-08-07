package main

import (
	"yannscrapy/logging"
	"yannscrapy/utils"
)


func main() {

	//var speed, status = ProxyTest("http://124.205.155.151:9090")
	//var speed, status = ProxyTest("http://05.252.161.48:8080")
	var speed, status = utils.DefaultCheckProxyIp("http://175.143.37.162:80")


	if status == 200 {
		logging.Infof("%d %d", speed, status)
	} else {
		logging.Info("代理不可用")
	}
}
