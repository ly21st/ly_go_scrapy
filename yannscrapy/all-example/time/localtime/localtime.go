package main

import (
	"fmt"
	"time"
)

func main() {

	layout := "2006-01-02 15:04:05"

	local, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation(layout, "2018-08-16 19:36:40", local)

	//t1, _ := time.ParseInLocation(layout, "2016-12-04 15:39:06 +0800 CST", local)


	fmt.Println("no.unix=", time.Now().UTC().Unix(), "\nt1.utc =", t1.UTC().Unix(), "\nt1.unix=", t1.Unix())

	fmt.Printf("t1=%v\n", t1)

	now := time.Now()
	fmt.Printf("now=%v\n", now)


	t1, _ = time.ParseInLocation(layout, "2021-08-31 00:00:00", local)
	t2, _ := time.ParseInLocation(layout, "2021-09-17 00:00:00", local)

	fmt.Printf("t1=%v\n", t1)
	fmt.Printf("t2=%v\n", t2)


	//for {
	//	if true == TriggerCheckSsp() {
	//		fmt.Println("----!!!!!---")
	//		break
	//	}
	//}

}

func TriggerCheckSsp() bool {
	fmt.Println("----")
	layout := "2006-01-02 15:04:05"

	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(layout, "2018-08-17 7:49:50", local)

	fmt.Println("---", time.Now().UTC().Unix(), "--", t.UTC().Unix(), "---", time.Now().Unix())
	if time.Now().UTC().Unix() >= t.UTC().Unix() {
		return true
	}
	return false
}