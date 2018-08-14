package service

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"context"
	"github.com/cyjme/service-center/client"
)

func ListNodeByServiceName(serviceName string) {

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Get(context.TODO(), Prefix+serviceName, clientv3.WithPrefix())

	if err != nil {
		log.Println("list node by serviceName", err)
	}
	log.Println("list node by serviceName", result.Kvs)
}
