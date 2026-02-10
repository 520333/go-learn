package loadbalance

import "errors"

// RoundRobinBalance 轮询算法
type RoundRobinBalance struct {
	servAddr []string // 服务器主机地址 IP:host
	curIndex int      // 当前轮询的节点索引
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params langth at least 1")
	}
	for i := 0; i < len(params); i++ {
		r.servAddr = append(r.servAddr, params[i])
	}
	return nil
}

func (r *RoundRobinBalance) Next() string {
	lens := len(r.servAddr)
	if lens == 0 {
		return ""
	}
	if r.curIndex >= lens {
		r.curIndex = 0
	}
	addr := r.servAddr[r.curIndex]
	//r.curIndex++
	r.curIndex = (r.curIndex + 1) % lens
	return addr
}
