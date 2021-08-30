package main

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"github.com/clbanning/mxj"
)

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type Base struct {
	FromUserName CDATAText
	ToUserName   CDATAText
	MsgType      CDATAText
	CreateTime   CDATAText
}

type TextMessage struct {
	XMLName xml.Name `xml:"xml"`
	Base
	Content CDATAText
}

func value2CDATA(v string) CDATAText {
	return CDATAText{"<![CDATA[" + v + "]]>"}
}

func main() {
	// 1. 解析 XML
	xmlStr := `<xml> 
		<ToUserName><![CDATA[toUser]]></ToUserName> 
		<FromUserName><![CDATA[fromUser]]></FromUserName> 
		<CreateTime>1348831860</CreateTime> 
		<MsgType><![CDATA[text]]></MsgType> 
		<Content><![CDATA[this is a test]]></Content> 
		<MsgId>1234567890123456</MsgId> 
		</xml>`

	var Message map[string]interface{}
	m, err := mxj.NewMapXml([]byte(xmlStr))

	if err != nil {
		return
	}

	if _, ok := m["xml"]; !ok {
		fmt.Println("Invalid Message.")
		return
	}

	message, ok := m["xml"].(map[string]interface{}) // 转换类型

	if !ok {
		fmt.Println("Invalid Field `xml` Type.")
		return
	}

	Message = message

	fmt.Println("1. 解析出来:", Message)

	// 2. 封装XML
	var reply TextMessage
	inMsg, ok := Message["Content"].(string) // 读取内容

	if !ok {
		return
	}
	reply.Base.FromUserName = value2CDATA(Message["ToUserName"].(string))
	reply.Base.ToUserName = value2CDATA(Message["FromUserName"].(string))
	reply.Base.CreateTime = value2CDATA(strconv.FormatInt(time.Now().Unix(), 10))
	reply.Base.MsgType = value2CDATA("text")
	reply.Content = value2CDATA(fmt.Sprintf("我收到的是：%s", inMsg))

	replyXml, err := xml.Marshal(reply) // 得到的是byte
	fmt.Println("2. 生成XML:", string(replyXml))

}
