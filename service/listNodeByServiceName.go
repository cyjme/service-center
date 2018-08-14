package service

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"context"
	"github.com/cyjme/service-center/client"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

func ListNodeByServiceName(serviceName string) []*mvccpb.KeyValue {

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Get(context.TODO(), Prefix+serviceName, clientv3.WithPrefix())

	if err != nil {
		log.Println("list node by serviceName", err)
	}

	return result.Kvs
}
