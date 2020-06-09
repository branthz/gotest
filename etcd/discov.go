package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/branthz/utarrow/lib/log"

	"github.com/branthz/etcd/clientv3"
)

type Master struct {
	Path   string
	Nodes  map[string]*Node
	Client *clientv3.Client
}

//node is a client
type Node struct {
	State bool
	Key   string
	Info  map[uint64]*ServiceInfo
}

func NewMaster(endpoint string, watchPath string) (*Master, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Second,
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	master := &Master{
		Path:   watchPath,
		Nodes:  make(map[string]*Node),
		Client: cli,
	}

	go master.WatchNodes()
	return master, err
}

func (m *Master) show() {
	fmt.Println("lalalalalaal")
	for k, v := range m.Nodes {
		fmt.Printf("---%s\n", k)
		for kk, _ := range v.Info {
			fmt.Printf("+++%d\n", kk)
		}
	}
}
func (m *Master) AddNode(key string, info *ServiceInfo) {
	node, ok := m.Nodes[key]
	if !ok {
		node = &Node{
			State: true,
			Key:   key,
			Info:  map[uint64]*ServiceInfo{info.ID: info},
		}
		m.Nodes[node.Key] = node
	} else {
		node.Info[info.ID] = info
	}
}

func (m *Master) DeleteNode(key string, info *ServiceInfo) {
	node, ok := m.Nodes[key]
	if !ok {
		return
	} else {
		delete(node.Info, info.ID)
	}
}

func GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Errorln(err)
	}
	return info
}

func (m *Master) WatchNodes() {
	rch := m.Client.Watch(context.Background(), m.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				log.Info("[%s] %q : %q", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := GetServiceInfo(ev)
				m.AddNode(string(ev.Kv.Key), info)
			case clientv3.EventTypeDelete:
				log.Info("[%s] %v : %v", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := GetServiceInfo(ev)
				m.DeleteNode(string(ev.Kv.Key), info)
			}
		}
	}
}
