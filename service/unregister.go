package service

import (
	"golang.org/x/net/context"

	"github.com/coreos/etcd/clientv3"
	"service-center/client"
)

func Unregister(key string) {
	key = Prefix + key + "/"

	kv := clientv3.NewKV(client.Client)
	kv.Delete(context.TODO(), key)
}
