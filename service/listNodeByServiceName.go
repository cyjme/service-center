package service

import (
	"context"
	"log"
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/cyjme/service-center/client"
)

type Service struct {
	Name  string `json:"key"`
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Key string `json:"key"`
	Url string `json:"url"`
}

func ListNodeByServiceName(serviceName string) map[string][]Node {

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Get(context.TODO(), Prefix+serviceName, clientv3.WithPrefix())

	if err != nil {
		log.Println("list node by serviceName", err)
	}
	serviceMap := map[string][]Node{}

	for _, kv := range result.Kvs {
		key := string(kv.Key[:])
		value := string(kv.Value[:])

		keyArr := strings.Split(key, "/")
		serviceName := keyArr[1]
		nodeKey := keyArr[2]
		node := Node{
			Key: nodeKey,
			Url: value,
		}
		if serviceMap[serviceName] == nil {
			serviceMap[serviceName] = []Node{}
		}

		serviceMap[serviceName] = append(serviceMap[serviceName], node)
	}

	return serviceMap
}
