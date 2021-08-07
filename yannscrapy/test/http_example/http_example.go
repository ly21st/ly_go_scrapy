package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"yannscrapy/logging"
	"yannscrapy/utils"
)

func main() {
	httpRequest := new(utils.HttpRequest)
	httpRequest.Timeout = 60
	//url := "https://www.anadf.com/"
	url := "https://www.anadf.com/cn/"
	response, err := httpRequest.Request(http.MethodGet, url)
	if err != nil {
		logging.Fatalf("request failed:%v\n", err)
		os.Exit(1)
	}

	dstFileName := "aa.html"
	ioutil.WriteFile(dstFileName, response.Body(), 0600)
	logging.Info("finish")
}
