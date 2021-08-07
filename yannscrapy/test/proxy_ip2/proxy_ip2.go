package main

import (
	"yannscrapy/logging"
	"yannscrapy/utils"
)

func main() {
	url := "https://ip.jiangxianli.com/api/proxy_ips"
	ipList, err := utils.GetProxyips(url)
	if err != nil {
		logging.Error(err)
	}
	logging.Infof("ipList=%v", ipList)
	logging.Info("finish")
}


// 获取到的数据
//var example=`
//	{
//    "code":0,
//    "msg":"\u6210\u529f",
//    "data":{
//        "current_page":1,
//        "data":[
//            {
//                "unique_id":"38a8731be83aa0150b08c96e15c1cf2a",
//                "ip":"05.252.161.48",
//                "port":"8080",
//                "country":"\u4e2d\u56fd",
//                "ip_address":"\u4e2d\u56fd \u5c71\u4e1c\u7701 \u5411\u9633",
//                "anonymity":1,
//                "protocol":"http",
//                "isp":"Chinanet",
//                "speed":1130,
//                "validated_at":"2021-08-07 16:55:29",
//                "created_at":"2021-08-02 08:50:10",
//                "updated_at":"2021-08-07 16:55:29"
//            }
//        ]
//    }
//}
//`























