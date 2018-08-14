package test

import (
	"github.com/cyjme/service-center/client"
	"github.com/cyjme/service-center/config"
	"github.com/cyjme/service-center/service"
	"fmt"
	"time"
	"testing"
)

func Test_all(t *testing.T) {
	client.Init("localhost:2379")
	config.Set("user-center/db/host", "127.0.0.3")
	config.Set("user-center/db/name", "127.0.0.4")
	config.Set("user-center/db", "127.0.0.4")
	fmt.Println("dbname", config.Get("user-center/db/name"))

	go service.Register("user-center", "0.0.0.0:10001")
	go service.Register("user-center", "0.0.0.0:10002")
	go service.Register("user-center", "0.0.0.0:10003")
	go service.Register("user-center", "0.0.0.0:10004")

	time.Sleep(time.Second * 1)
	kvs := service.ListNodeByServiceName("user-center")

	fmt.Println("all key value", kvs)
}
