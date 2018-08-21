package service

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"context"
	"github.com/cyjme/service-center/client"
)

type Service struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ListNodeByServiceName(serviceName string) []Service {

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Get(context.TODO(), Prefix+serviceName, clientv3.WithPrefix())

	if err != nil {
		log.Println("list node by serviceName", err)
	}

	serviceArray := []Service{}

	for _, kv := range result.Kvs {
		serviceArray = append(serviceArray, Service{
			Key:   string(kv.Key[:]),
			Value: string(kv.Value[:]),
		})
	}

	return serviceArray
}
