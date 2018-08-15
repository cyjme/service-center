package service

import (
	"github.com/cyjme/service-center/client"
	"context"
	"github.com/coreos/etcd/clientv3"
	"fmt"
	"log"
)

func Watch(serviceName string, callback func()) {
	responseWatchChannel := client.Client.Watch(context.Background(), serviceName, clientv3.WithPrefix())

	for wresp := range responseWatchChannel {
		for _, ev := range wresp.Events {
			log.Println("watch change")
			log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

			callback()
		}
	}
}
