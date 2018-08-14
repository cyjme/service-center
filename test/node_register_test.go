package test

import (
	"service-center/client"
	"service-center/config"
	"service-center/service"
	"fmt"
	"time"
	"testing"
)

func Test_all(t *testing.T) {
	client.Init()
	config.Set("user-center/db/host", "127.0.0.3")
	config.Set("user-center/db/name", "127.0.0.4")
	config.Set("user-center/db", "127.0.0.4")
	config.Get("user-center")

	go service.Register("user-center", "0.0.0.0:10001")
	go service.Register("user-center", "0.0.0.0:10002")
	go service.Register("user-center", "0.0.0.0:10003")
	go service.Register("user-center", "0.0.0.0:10004")

	time.Sleep(time.Second * 1)
	service.ListNodeByServiceName("user-center")
	fmt.Println("all key value")
}
