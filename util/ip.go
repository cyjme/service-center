package util

import (
	"net"
	"net/http"
	"os"
	"log"
	"errors"
	"io/ioutil"
)

//GetIntranetIp 获取内网 ip
func GetIntranetIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("get intranetIp error")
}

//GetPublicIp 获取公网 ip
func GetPublicIp() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	return string(content), err
}
