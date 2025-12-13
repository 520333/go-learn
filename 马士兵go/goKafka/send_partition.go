package goKafka

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

func SendPartition() {
	// 1.得到异步的producer
	broker := []string{"192.168.50.100:9092"}
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true //开启success channel来接受发送成功的信息

	// 配置producer选项分区的策略
	//conf.Producer.Partitioner = sarama.NewRandomPartitioner //随机 不需要逻辑上的分割
	//conf.Producer.Partitioner = sarama.NewRandomPartitioner //分区调度策略轮询
	//conf.Producer.Partitioner = sarama.NewHashPartitioner   //Hash调度策略 基于特定的key
	conf.Producer.Partitioner = sarama.NewManualPartitioner //指定分区调度策略 配合partition使用
	//conf.Producer.Partitioner = sarama.NewCustomHashPartitioner(myfunc() hash.Hash32 {
	//	return crc32.New(crc32.MakeTable(0xD5828281))
	//}) //自定义分区HASH调度策略

	producer, err := sarama.NewAsyncProducer(broker, conf)
	if err != nil {
		log.Fatalf("Error creating producer: %s\n", err)
	}

	// 2.启用goroutine
	var wg sync.WaitGroup
	var sendCounter, successCounter, errorCount = 0, 0, 0

	// 2.1处理errors
	go func() {
		wg.Add(1)
		defer wg.Done()
		for err := range producer.Errors() {
			log.Printf("Failed to send err: %v\n", err)
			errorCount++
		}
	}()
	// 2.2处理success
	go func() {
		wg.Add(1)
		defer wg.Done()
		for suc := range producer.Successes() {
			log.Printf("Success to send, partition:%v , offset:%v, value:%v\n", suc.Partition, suc.Offset, suc.Value)
			successCounter++
		}
	}()
	// 发送消息
	// 增加控制信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt) //ctrl + c
loop:
	for {
		// 生成特定的消息
		var id = rand.Intn(5)
		fmt.Println("current message id:", id)
		time.Sleep(500 * time.Millisecond)
		select {
		case producer.Input() <- &sarama.ProducerMessage{
			Topic:     "topic_more_partition_1",
			Value:     sarama.StringEncoder("海绵宝宝 Go,id:" + fmt.Sprintf("%d", id)),
			Key:       sarama.StringEncoder(fmt.Sprintf("%d", id)),
			Partition: int32(id % 3),
		}:
			sendCounter++
		case <-signals:
			producer.AsyncClose() //异步终止
			break loop
		}
	}
	log.Printf("Send counter: %v, error count: %v SuccessCounter: %v,\n", sendCounter, errorCount, successCounter)

	wg.Wait()

}
