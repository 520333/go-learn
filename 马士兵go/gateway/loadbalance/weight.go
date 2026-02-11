package loadbalance

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	maxFails    = 3               // 最大失败次数
	failTimeout = time.Second * 3 // 失败次数最多之后
)

// WeightRoundRobinBalance 加权轮询
type WeightRoundRobinBalance struct {
	servAddr []*node // 服务器主机地址 IP:host
	curIndex int     // 当前轮询的节点索引
}

// node 每一个服务器节点有不同的权重 并且在每一轮访问后可能会发生变化
type node struct {
	addr            string        // 主机地址
	weight          int           // 初始化权重
	currentWeight   int           // 节点当前临时权限
	effectiveWeight int           // 有效权重
	maxFails        int           // 在failTimeout时间内最大的失败次数
	failTimeout     time.Duration // 超时时间 衡量最大失败次数
	failTimes       []time.Time   // 失败的时间点 按时间正序排列
}

// Add 添加带权重的服务器主机
func (r *WeightRoundRobinBalance) Add(params ...string) error {
	length := len(params)
	if length == 0 || length%2 != 0 {
		return errors.New("param's length must be 2 or multiple of 2")
	}
	for i := 0; i < length; i += 2 {
		addr := params[i]
		weight, err := strconv.ParseInt(params[i+1], 10, 32)
		if err != nil {
			return err
		}
		if weight < 0 {
			weight = 1
		}
		n := &node{
			addr:            addr,
			weight:          int(weight),
			effectiveWeight: int(weight),
			maxFails:        maxFails,
			failTimeout:     failTimeout,
		}
		r.servAddr = append(r.servAddr, n)
	}
	return nil
}

// Next 获取下一个服务器地址，找到权重最大的服务器
func (r *WeightRoundRobinBalance) Next() (string, error) {
	var index = 0
	var effectiveTotal = 0
	var maxNode *node
	for i := 0; i < len(r.servAddr); i++ {
		w := r.servAddr[i]
		if w.maxFails <= 0 {
			refreshErrRecords(w)
			w.maxFails = maxFails - len(w.failTimes)
			if w.maxFails <= 0 {
				fmt.Println("小黑屋：" + w.addr)
				//time.Sleep(time.Second * 3)
				continue
			}
		}
		w.currentWeight += w.effectiveWeight
		if maxNode == nil || w.currentWeight > maxNode.currentWeight {
			maxNode = w
			index = i
		}
		effectiveTotal += w.effectiveWeight
	}
	if maxNode == nil {
		// 服务器列表为空 返回error
		return "", errors.New("there is no server address. please call Add(...string)")
	}
	maxNode.currentWeight -= effectiveTotal
	r.curIndex = index
	return maxNode.addr, nil
}

func (r *WeightRoundRobinBalance) callback(addr string, flag bool) {
	for i := 0; i < len(r.servAddr); i++ {
		w := r.servAddr[i]
		if w.addr == addr {
			// 访问服务器成功
			if flag {
				if w.effectiveWeight < w.weight {
					w.effectiveWeight++
				}
			}
		} else {
			// 访问服务器失败
			w.effectiveWeight--
			// 刷新错误记录
			refreshErrRecords(w)
			w.failTimes = append(w.failTimes, time.Now())
			w.maxFails = maxFails - len(w.failTimes)
		}
		break
	}
}

// refreshErrRecords 刷新错误记录，把过去超过failTimeout 的错误记录删除
func refreshErrRecords(w *node) {
	now := time.Now()
	var i = 0
	for ; i < len(w.failTimes); i++ {
		failTime := w.failTimes[i].Add(failTimeout)
		if failTime.After(now) || failTime.Equal(now) {
			break
		}
	}
	w.failTimes = w.failTimes[i:]
}
