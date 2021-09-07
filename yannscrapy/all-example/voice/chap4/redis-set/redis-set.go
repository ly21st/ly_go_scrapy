package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	c, err := redis.Dial("tcp", "8.134.34.59:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	log.Println("connect redis ok")

	defer c.Close()
	_, err = c.Do("Set", "count", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("set ok")

	r, err := redis.Int(c.Do("Get", "count"))
	if err != nil {
		fmt.Println("get count failed,", err)
		return
	}
	log.Printf("get ok")
	fmt.Println(r)
}
