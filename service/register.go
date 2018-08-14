package service

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"service-center/client"
)

func Register(serviceName string, value string) {
	key := serviceName
	key = Prefix + key + "/"

	kv := clientv3.NewKV(client.Client)
	lease := clientv3.NewLease(client.Client)
	var curLeaseId clientv3.LeaseID = 0
	for {
		if curLeaseId == 0 {
			leaseResp, err := lease.Grant(context.TODO(), 10)
			if err != nil {
				goto SLEEP
			}

			key := key + fmt.Sprintf("%d", leaseResp.ID)
			if _, err := kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseResp.ID)); err != nil {
				goto SLEEP
			}
			curLeaseId = leaseResp.ID
		} else {
			if _, err := lease.KeepAliveOnce(context.TODO(), curLeaseId); err == rpctypes.ErrLeaseNotFound {
				curLeaseId = 0
				continue
			}
		}
	SLEEP:
		time.Sleep(time.Duration(1) * time.Second)
	}
}
