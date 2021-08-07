package utils

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
	"yannscrapy/config"

	"gopkg.in/resty.v1"
)

type HttpRequest struct {
	Timeout time.Duration
}

const (
	DefaultTimeout = 10
)

// 获取ip地址
func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("cannot connect to the network")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil
	}

	return ip
}

//  获取ip地址
func GetServerIp() string {
	var ip string
	if config.Conf.Address == "" || config.Conf.Address == "0.0.0.0" {
		ip, _ = ExternalIP()
	} else {
		ip = config.Conf.Address
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	return ip
}

// http rest request encapsulated simply
// e.g: request(method, url, params, header)
func (httpRequest *HttpRequest) Request(method, url string, args ...interface{}) (*resty.Response, error) {
	var params interface{}
	header := map[string]string{}
	switch len(args) {
	case 1:
		params = args[0]
	case 2:
		header = args[1].(map[string]string)
	}
	header["content-type"] = "application/json"

	var res *resty.Response
	var err error
	if httpRequest.Timeout == 0 {
		httpRequest.Timeout = DefaultTimeout
	}
	req := resty.New().SetTimeout(time.Second * httpRequest.Timeout).R().SetHeaders(header)
	switch method {
	case http.MethodGet:
		res, err = req.Get(url)
	case http.MethodPost:
		res, err = req.SetBody(params).Post(url)
	case http.MethodPut:
		res, err = req.SetBody(params).Put(url)
	case http.MethodDelete:
		res, err = req.Delete(url)
	default:
		return nil, errors.New("Failed to recognize method: " + method)
	}
	if err != nil {
		return nil, err
	}
	if res.StatusCode() >= 400 {
		return nil, errors.New(fmt.Sprintf("Request failed, Http status code: %d, body: %s", res.StatusCode(), string(res.Body())))
	}

	return res, err
}
