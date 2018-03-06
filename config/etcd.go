package config

import (
	"context"
	"strings"

	"github.com/coreos/etcd/clientv3"
)

type Receiver struct {
	evType string
	key    string
	value  string
}

var ConfigChan = make(chan *Receiver)

func InitEtcdConifg() error {
	if UseEtcd == false {
		return nil
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(Default.Etcd.Endpoints, ","),
	})
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := client.Get(context.TODO(), "address", clientv3.WithPrefix())
	if err != nil {
		return err
	}

	Default.Base.Address = string(resp.Kvs[0].Value)

	return nil
}
