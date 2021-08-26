package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func main() {

	origin := "http://localhost/"
	url := "ws://localhost:8972/api/v1/goods/log"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := ws.Write([]byte("hello world!")); err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	var n int


		//if n, err = ws.Read(msg); err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("Received: %s.\n", msg[:n])
		//time.Sleep(1 * time.Second)



	for {
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received: %s.\n", msg[:n])
		time.Sleep(1 * time.Second)
	}
}

func time_test() {
	//fmt.Printf(time.Now().String())

	now  := time.Now()
	//Year = now.Year()
	//Mouth  = now.Month()
	//Day  =  now.Day()
	//时间格式化输出 Printf输出
	fmt.Printf("当前时间为： %d-%d-%d %d:%d:%d\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	//fmt.Sprintf 格式化输出
	dateString := fmt.Sprintf("当前时间为： %d-%d-%d %d:%d:%d\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(dateString)
	//now.Format 方法格式化
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02"))//年月日
	fmt.Println(now.Format("15:04:05"))//时分秒
}