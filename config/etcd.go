package config

import (
	"context"
	"reflect"
	"strings"

	"go.etcd.io/etcd/clientv3"
)

//InitConifgFromEtcd 从etcd初始化配置信息
func InitConifgFromEtcd() error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: strings.Split(Default.Etcd.Endpoints, ","),
	})
	if err != nil {
		return err
	}
	defer client.Close()
	//base
	v := reflect.ValueOf(Default.Base)
	t := reflect.TypeOf(Default.Base)
	m := map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("json")
		str := getValueFromEtcd(client, "config/base/"+fieldName)
		m[fieldName] = str
	}
	setStructFieldByJSONNamego(&Default.Base, m)

	//mongodb
	v = reflect.ValueOf(Default.MongoDB)
	t = reflect.TypeOf(Default.MongoDB)
	m = map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("json")
		str := getValueFromEtcd(client, "config/mongodb/"+fieldName)
		m[fieldName] = str
	}
	setStructFieldByJSONNamego(&Default.MongoDB, m)

	//session
	v = reflect.ValueOf(Default.Session)
	t = reflect.TypeOf(Default.Session)
	m = map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Tag.Get("json")
		str := getValueFromEtcd(client, "config/session/"+fieldName)
		m[fieldName] = str
	}
	setStructFieldByJSONNamego(&Default.Session, m)

	return nil
}

func getValueFromEtcd(client *clientv3.Client, key string) string {
	resp, err := client.Get(context.TODO(), key, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	if len(resp.Kvs) == 0 {
		return ""
	}

	return string(resp.Kvs[0].Value)
}

func setStructFieldByJSONNamego(structPtr interface{}, fields map[string]interface{}) {
	cType := reflect.TypeOf(structPtr)
	cValue := reflect.ValueOf(structPtr).Elem()
	structLen := cValue.NumField()
	for i := 0; i < structLen; i++ {
		field := cType.Elem().Field(i)
		jsonName := field.Tag.Get("json")
		if jsonName == "" {
			continue
		}
		//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
		jsonName = strings.Split(jsonName, ",")[0]
		if value, ok := fields[jsonName]; ok {
			//给结构体赋值
			if reflect.ValueOf(value).Type() == cValue.FieldByName(field.Name).Type() {
				cValue.FieldByName(field.Name).Set(reflect.ValueOf(value))
			}

		}
	}
}
