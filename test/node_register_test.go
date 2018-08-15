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
	config.Set("user-center", `
http:
  domain: "0.0.0.0"
  port: 10001
rpc:
  domain: "0.0.0.0"
  port: 10002
db:
  host: "127.0.0.1"
  port: "3306"
  user: "root"
  password: ""
  name: "user-center"
redis:
  addr: "127.0.0.1:6379"
  password: ""
  pool_size: 10
  db: 11
secret:
  jwt_key: "jwt_key"
  pass_hash_salt: "pass_hash_salt"
sendcloud:
  api_key: "3dmnlw2ebuO4GiOuV1iWyI4N3NVOUD61"
  sms_user: "ipar"
  msg_type: "0"
`)
	//config.Set("user-center/db/host", "127.0.0.3")
	//config.Set("user-center/db/name", "127.0.0.4")
	//config.Set("user-center/db", "127.0.0.4")
	//fmt.Println("dbname", config.Get("user-center/db/name"))

	//go service.Register("user-center", "0.0.0.0:10001")
	//go service.Register("user-center", "0.0.0.0:10002")
	//go service.Register("user-center", "0.0.0.0:10003")
	//go service.Register("user-center", "0.0.0.0:10004")

	time.Sleep(time.Second * 1)
	kvs := service.ListNodeByServiceName("user-center")

	fmt.Println("all key value", kvs)
}
