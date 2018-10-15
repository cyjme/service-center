package config

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/coreos/etcd/clientv3"
	"github.com/cyjme/service-center/client"
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

func List(key string) map[string]string {
	key = Prefix + key

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Get(context.TODO(), key, clientv3.WithPrefix())

	if err != nil {
		log.Println("get config error", err)
	}

	configList := map[string]string{}
	for _, kv := range result.Kvs {
		key := string(kv.Key[:])
		value := string(kv.Value[:])
		//delete key prefix config/
		index := strings.Index(key, "/")
		key = key[index+1:]

		configList[key] = string(value)
	}

	return configList
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

func Clear(key string) {
	key = Prefix + key

	kv := clientv3.NewKV(client.Client)
	result, err := kv.Delete(context.TODO(), key)

	if err != nil {
		log.Println("clear config error", err)
	}
	log.Println("clear config result", result)
}

func ConvertPathToMap(path string, resultMap *map[string]interface{}) {
	// keyStringArr := strings.SplitN(path, "/", 1)
	index := strings.Index(path, "/")
	if index == -1 {
		(*resultMap)[path] = make(map[string]interface{}, 0)
		return
	}

	fmt.Println("index", index)
	fmt.Println("str", path[:index])

	newPath := path[index+1:]
	fmt.Println("newPath", newPath)

	//todo?
	//获取 newPath 中的第一个字段

	if (*resultMap)[path] == nil {
		(*resultMap)[path] = make(map[string]interface{}, 0)
	}
}

func getBeforeKey(str string, key string) {
}
