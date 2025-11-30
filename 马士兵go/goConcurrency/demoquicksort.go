package goConcurrency

import (
	"math/rand"
	"sync"
	"time"
)

func QuickSortConcurrency(arr []int) []int {
	// 1.校验arr是否满足排序需要 至少2个元素
	if arr == nil || len(arr) < 2 {
		return arr
	}
	// 4.同步控制
	wg := &sync.WaitGroup{}
	// 2.执行排序
	wg.Add(1)
	go quickSortConcurrency(arr, 0, len(arr)-1, wg)
	wg.Wait()

	// 3.返回值
	return arr
}

// 实现递归排序的核心函数
func quickSortConcurrency(arr []int, l, r int, wg *sync.WaitGroup) {
	// 一：-1wg计数器
	defer wg.Done()
	// 二：是否需要排序 l<r
	if l < r {
		// 三：大小分区元素，并获取参考袁术索引
		mid := partition(arr, l, r)
		// 四：并发对左部分排序
		wg.Add(1)
		go quickSortConcurrency(arr, l, mid-1, wg)

		// 五：并发的对右部分排序
		wg.Add(1)
		go quickSortConcurrency(arr, mid+1, r, wg)
	}
}
func partition(arr []int, l, r int) int {
	p := l - 1
	for i := l; i <= r; i++ {
		if arr[i] <= arr[r] {
			p++
			swap(arr, p, i)
		}
	}
	return p
}
func swap(arr []int, i, j int) {
	t := arr[i]
	arr[i] = arr[j]
	arr[j] = t
}

func GenerateRandArr(l int) []int {
	arr := make([]int, l)
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < l; i++ {
		arr[i] = int(rand.Int31n(int32(l * 5)))
	}
	return arr
}
