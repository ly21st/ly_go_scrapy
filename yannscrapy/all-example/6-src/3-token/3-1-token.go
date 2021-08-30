package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"time"
)

func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

// 根据 Token Timestamp Nonce 生成对应的校验码， Token是不能明文传输的
func GenerateSignature(token string) (timestamp string, nonce string, signature string) {

	nonce = CreateRandomString(10)
	timestamp = strconv.FormatInt(time.Now().Unix(), 10)

	strs := sort.StringSlice{token, timestamp, nonce} // 使用本地的token生成校验
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func VerifySignature(token string, timestamp string, nonce string, signature string) bool {
	strs := sort.StringSlice{token, timestamp, nonce} // 使用本地的token生成校验
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil)) == signature
}
func main() {
	token := "liaoqingfu"
	// 产生签名
	timestamp, nonce, signature := GenerateSignature(token) // 发送服务器的时候是发送  timestamp, nonce, signature
	fmt.Printf("1. token %s -> 产生签名:%s\n", token, signature)
	// 验证签名
	ok := VerifySignature(token, timestamp, nonce, signature)
	if ok {
		fmt.Println("2. 验证签名正常")
	} else {
		fmt.Println("2. 验证签名失败")
	}
}
