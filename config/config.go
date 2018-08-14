package config

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/cyjme/service-center/client"
	"log"
	"context"
)

const Prefix = "config/"

func Set(key string, value string) {
	key = Prefix + key

	kv := clientv3.NewKV(client.Client)
	_, err := kv.Put(context.TODO(), key, value)

	if err != nil {
		log.Println("set config error", err)
	}
	log.Println("set config success", key, value)
}

func Get(key string) string {
	key = Prefix + key

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Get(context.TODO(), key, clientv3.WithPrefix())

	if err != nil {
		log.Println("get config error", err)
	}

	return string(result.Kvs[0].Value[:])
}

func Clear() {
	key := Prefix

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Get(context.TODO(), key, clientv3.WithPrefix())

	if err != nil {
		log.Println("clear config error", err)
	}
	log.Println("clear config result", result.Kvs)
}