package loadbalance

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

// ConsistentHashBalance 一致性hash算法
type ConsistentHashBalance struct {
	hash     Hash              // hash函数 支持用户自定义 默认使用crc32.ChecksumIEEE
	hasKeys  Uint32Slice       // 服务器节点hash值 按照从小到大排序
	hashMap  map[uint32]string // 服务器节点的hash值与服务器真实地址的映射表
	replicas int               // 虚拟节点倍数
	mux      sync.RWMutex
}

type Hash func(data []byte) uint32

type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}

func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func NewConsistentHashBalance(replicas int, fn Hash) *ConsistentHashBalance {
	ch := &ConsistentHashBalance{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[uint32]string),
	}
	if ch.hash == nil {
		ch.hash = crc32.ChecksumIEEE
	}
	return ch
}

func (c *ConsistentHashBalance) Add(servers ...string) error {
	if len(servers) == 0 {
		return errors.New("servers length at least 1")
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	for _, addr := range servers {
		for i := 0; i < c.replicas; i++ {
			hash := c.hash([]byte(strconv.Itoa(i) + addr))
			c.hasKeys = append(c.hasKeys, hash)
			c.hashMap[hash] = addr
		}
	}
	sort.Sort(c.hasKeys) // 对所有节点的hash值进行排序
	return nil
}

func (c *ConsistentHashBalance) Get(key string) (string, error) {
	l := len(c.hasKeys)
	if l == 0 {
		return "", errors.New("node list is empty")
	}
	hash := c.hash([]byte(key))
	fmt.Println(strconv.FormatInt(int64(hash), 10) + ":")
	index := sort.Search(l, func(i int) bool {
		return c.hasKeys[i] >= hash
	})
	if index == l {
		index = 0
	}
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.hashMap[c.hasKeys[index]], nil
}
