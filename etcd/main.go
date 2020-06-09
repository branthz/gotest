package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/branthz/etcd/clientv3"
	"github.com/branthz/utarrow/lib/log"
)

type ServiceInfo struct {
	ID uint64
	IP string
}
type Service struct {
	Name    string
	client  *clientv3.Client
	Info    ServiceInfo
	leaseid clientv3.LeaseID
	stop    chan error
}

func init() {
	log.Setup("", "debug")
}

func (s *Service) start() error {
	ch, err := s.keepAlive()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	for {
		select {
		case err := <-s.stop:
			s.revoke()
			return err
		case <-s.client.Ctx().Done():
			return errors.New("server closed")
		case ka, ok := <-ch:
			if !ok {
				log.Infoln("keep alive channel closed")
				s.revoke()
				return nil
			} else {
				log.Info("Recv reply from service: %s, ttl:%d", s.Name, ka.TTL)
			}
		}
	}
}
func (s *Service) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	resp, err := s.client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatalln(err)
	}
	key := "services/" + s.Name
	value, _ := json.Marshal(s.Info)
	_, err = s.client.Put(context.TODO(), key, string(value), clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	s.leaseid = resp.ID
	return s.client.KeepAlive(context.TODO(), resp.ID)
}

func newService(name, ip, etcd string, id uint64) (*Service, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcd},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &Service{
		Name:   name,
		client: cli,
		Info:   ServiceInfo{ID: id, IP: ip},
		stop:   make(chan error),
	}, nil
}
func (s *Service) Stop() {
	s.stop <- nil
}
func (s *Service) revoke() error {
	_, err := s.client.Revoke(context.TODO(), s.leaseid)
	if err != nil {
		log.Fatalln(err)
	}
	log.Info("servide:%s stop\n", s.Name)
	return err
}

func main() {
	s, _ := newService("controller", "1.1.1.1", "127.0.0.1:2379", 100)
	go s.start()
	m, err := NewMaster("127.0.0.1:2379", "services/controller")
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1e9)
	m.show()
	s.Stop()
	time.Sleep(6 * 1e9)
	m.show()
	time.Sleep(1e9)
}
