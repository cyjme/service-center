package client

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"log"
)

var Client *clientv3.Client

func Init(url string) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{url},
		DialTimeout: time.Second * 5,
	})

	if err != nil {
		log.Fatal("error", err)
	}

	Client = client

	log.Println("client init success")
}
