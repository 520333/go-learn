package loadbalance

import (
	"errors"
	"math/rand"
)

// RandomBalance 随机算法
type RandomBalance struct {
	servAddr []string // 服务器主机地址 IP:host
	curIndex int      // 当前轮询的节点索引
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params length at least 1")
	}
	for i := 0; i < len(params); i++ {
		r.servAddr = append(r.servAddr, params[i])
	}
	return nil
}

func (r *RandomBalance) Next() string {
	lens := len(r.servAddr)
	if lens == 0 {
		return ""
	}
	if r.curIndex >= lens {
		r.curIndex = 0
	}
	r.curIndex = rand.Intn(len(r.servAddr))
	return r.servAddr[r.curIndex]
}
