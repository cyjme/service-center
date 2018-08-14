package service

import (
	"github.com/cyjme/service-center/client"
	"context"
	"github.com/coreos/etcd/clientv3"
	"fmt"
)

func Watch(serviceName string) {
	responseWatchChannel := client.Client.Watch(context.Background(), serviceName, clientv3.WithPrefix())

	for wresp := range responseWatchChannel {
		for _, ev := range wresp.Events {
			fmt.Println("watch change")
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
