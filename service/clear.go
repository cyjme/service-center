package service

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"context"
	"service-center/client"
)

func Clear() {
	key := Prefix

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Delete(context.TODO(), key, clientv3.WithPrefix())

	if err != nil {
		log.Println("clear kv error", err)
	}
	log.Println("clear service result", result)
}
