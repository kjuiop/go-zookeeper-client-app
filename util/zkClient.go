package util

import (
	"fmt"
	"github.com/catenoid-company/curator.go"
	"time"
)

type ZKCon struct {
	zkClient  curator.CuratorFramework
	connected bool
	cfg       *Config
}

func ZookeeperInitialize(cfg *Config) (*ZKCon, error) {

	zkClient := &ZKCon{
		zkClient:  nil,
		connected: false,
		cfg:       cfg,
	}

	// 주키퍼 접속
	err := zkClient.StartZookeeper()
	if err != nil {
		return nil, err
	}

	return zkClient, nil
}

func (zkCon *ZKCon) StartZookeeper() error {

	root := zkCon.cfg.ZookeeperInfo.RootNode
	// 접속 timeout 정책 설정
	retryPolicy := curator.NewExponentialBackoffRetry(time.Second, 3, 15*time.Second)
	// 클라이언트를 생성한다.
	zkCon.zkClient = curator.NewClient(zkCon.cfg.ZookeeperInfo.Host, retryPolicy)
	err := zkCon.zkClient.Start()
	if err != nil {
		return err
	}

	// 연결이 완료될 때까지 대기한다.
	err = zkCon.zkClient.ZookeeperClient().BlockUntilConnectedOrTimedOut()
	if err != nil {
		return err
	}
	fmt.Sprintf("Zookeeper(%s) connected, OK", zkCon.cfg.ZookeeperInfo.Host)

	// 최상위 노드 존재 확인
	err = zkCon.CheckExists(root)
	if err != nil {
		return err
	}

	// jobmanager 노드 존재 확인
	err = zkCon.CheckExists(root + "/disable")
	if err != nil {
		return err
	}

	return nil
}

// CheckExists 해당 경로의 node가 존재하는지 확인하고 미 존재 시 생성 처리
func (zkCon *ZKCon) CheckExists(path string) (err error) {
	stat, err := zkCon.zkClient.CheckExists().ForPath(path)
	if err != nil {
		return err
	}

	// 노드가 존재하지 않으면 생성
	if stat == nil {
		_, err = zkCon.zkClient.Create().WithMode(curator.PERSISTENT).ForPath(path)
		if err != nil {
			return err
		}
	}

	return
}

// EndZookeeper zookeeper 접속 종료
func (zkCon *ZKCon) EndZookeeper() {
	if zkCon.zkClient != nil {
		zkCon.zkClient.Close()
	}
	zkCon.connected = false
}
