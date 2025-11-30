package goConcurrency

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// WalkDir 外部调用的遍历目录统计信息的方法
func WalkDir(dirs ...string) string {
	// 1.保证至少有一个目录需要统计遍历
	// 默认当前目录
	if len(dirs) == 0 {
		dirs = []string{"."}
	}
	// 2.初始化变量 channel用于完成size的传递，waitGroup用于等待调度
	filesizeCh := make(chan int64, 1)
	wg := &sync.WaitGroup{}

	// 3.启动多个Goroutine统计信息，取决于len(dirs)
	for _, dir := range dirs {
		wg.Add(1)
		go walkDir(dir, filesizeCh, wg)
	}
	// 4.启动累计运算的Goroutine
	// 用于关闭filesizeCh
	go func(wg *sync.WaitGroup) {
		wg.Wait() // 等待统计工作完成
		close(filesizeCh)
	}(wg)

	//统计结果通过chan传递出来
	fileNumCh := make(chan int64, 1)
	sizeTotalCh := make(chan int64, 1)
	go func(filesizeCh <-chan int64, fileNumCh, sizeTotalCh chan<- int64) {
		var fileNum, sizeTotal int64
		for filesize := range filesizeCh {
			// 累计文件数，和统计文件整体大小
			fileNum++
			sizeTotal += filesize
		}
		fileNumCh <- fileNum
		sizeTotalCh <- sizeTotal
	}(filesizeCh, fileNumCh, sizeTotalCh)
	// 5.整理返回值
	return fmt.Sprintf("%d files %.2f MB\n", <-fileNumCh, float64(<-sizeTotalCh)/(1024*1024)) //1e6
}

// 遍历并统计某个特定目录的信息 核心功能实现函数 完成递归 统计等
func walkDir(dir string, filesizeCh chan<- int64, wg *sync.WaitGroup) {
	// 1.wg计数器减少
	defer wg.Done()
	// 2.读取dir下的全部文件
	for _, fileinfo := range fileInfos(dir) {
		// 3.根据dir下的文件信息
		if fileinfo.IsDir() {
			// 3.1如果目录 递归获取文件
			subDir := filepath.Join(dir, fileinfo.Name())
			// 递归调用
			wg.Add(1)
			go walkDir(subDir, filesizeCh, wg)
		} else {
			// 3.2如果不是就是文件统计文件大小，放入channel
			filesizeCh <- fileinfo.Size()
		}
	}

}

// 获取目录下文件信息
func fileInfos(dir string) []fs.FileInfo {
	// 1.读取目录的全部文件
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Println("WarkDir error:", err)
		return []fs.FileInfo{}
	}
	// 2.获取文件的文件信息
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		if info, err := entry.Info(); err == nil {
			infos = append(infos, info)
		}

	}
	// 3.返回
	return infos
}
