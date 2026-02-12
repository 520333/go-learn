package loadbalance

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundRobin(t *testing.T) {
	rb := &RoundRobinBalance{}
	rb.Add("127.0.0.1:8001")
	rb.Add("127.0.0.1:8002")
	rb.Add("127.0.0.1:8003")
	rb.Add("127.0.0.1:8004")
	rb.Add("127.0.0.1:8005")

	for i := 0; i < 10; i++ {
		fmt.Println(rb.Next())
	}
}

func TestWeightRoundRobinBalance(t *testing.T) {
	rb := &WeightRoundRobinBalance{}
	rb.Add("127.0.0.1:8001", "6")
	rb.Add("127.0.0.1:8002", "3")
	rb.Add("127.0.0.1:8003", "1")
	print(rb, "")
	fmt.Println("---------- init over ----------")
	for i := 0; i < 15; i++ {
		addr, err := rb.Next()
		assert.Nil(t, err)

		var r = rand.Intn(2)
		if r == 1 { // 故障的概率
			fmt.Println("server" + addr + " has failed.")
			rb.callback(addr, false)
		} else {
			fmt.Println("正常访问：" + addr)
			rb.callback(addr, true)
		}
		print(rb, addr)
	}
}

func print(rb *WeightRoundRobinBalance, addr string) {
	fmt.Println("================================")
	fmt.Println("主机地址\t\t\t当前权重\t有效权重")
	total := 0
	for j := 0; j < len(rb.servAddr); j++ {
		w := rb.servAddr[j]
		total += w.effectiveWeight
		cw := strconv.Itoa(w.currentWeight)
		ew := strconv.Itoa(w.effectiveWeight)
		if w.addr == addr {
			fmt.Printf("%c[1;0;31m%s%c[0m", 0x1B, addr, 0x1B)
		} else {
			fmt.Print(w.addr)
		}
		var str = "\t\t" + cw + "\t\t" + ew + "\t\t"
		fmt.Println(str)
	}
	fmt.Println("有效权重之和: \t\t\t\t" + strconv.Itoa(total))
}

func TestRandomBalance(t *testing.T) {
	rb := &RandomBalance{}
	rb.Add("127.0.0.1:8001")
	rb.Add("127.0.0.1:8002")
	rb.Add("127.0.0.1:8003")
	rb.Add("127.0.0.1:8004")
	rb.Add("127.0.0.1:8005")

	for i := 0; i < 10; i++ {
		fmt.Println(rb.Next())
	}
}

func TestConsistentHashBalance(t *testing.T) {
	rb := NewConsistentHashBalance(2, nil)
	rb.Add("127.0.0.1:8001", "127.0.0.1:8002", "127.0.0.1:8003", "127.0.0.1:8004", "127.0.0.1:8005")
	//fmt.Println(rb.hasKeys)
	fmt.Println(rb.hashMap)
	funcName(rb)
	key := rb.hasKeys[2]
	rb.hasKeys = append(rb.hasKeys[:1], rb.hasKeys[2:]...)
	delete(rb.hashMap, key)
	funcName(rb)

}

func funcName(rb *ConsistentHashBalance) {
	fmt.Println("=============")
	fmt.Println(rb.Get("http://127.0.0.1:8002/demo/get"))
	fmt.Println(rb.Get("http://127.0.0.1:8003/demo/getDemo"))
	fmt.Println(rb.Get("http://127.0.0.1:8002/demo/get"))
	fmt.Println(rb.Get("http://127.0.0.1:8004/demo/getBalance"))
	fmt.Println("--------------")
	fmt.Println(rb.Get("127.0.0.1:8002"))
	fmt.Println(rb.Get("127.0.0.1:8003"))
	fmt.Println(rb.Get("127.0.0.1:8002"))
	fmt.Println(rb.Get("192.168.1.254:8004"))
}
