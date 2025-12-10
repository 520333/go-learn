package goKafka

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

func AsyncSend() {
	// 1.得到异步的producer
	broker := []string{"192.168.50.100:9092"}
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true //开启success channel来接受发送成功的信息
	producer, err := sarama.NewAsyncProducer(broker, conf)
	if err != nil {
		log.Fatalf("Error creating producer: %s\n", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing producer: %s\n", err)
		}
	}()
	// 增加控制信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt) //ctrl + c

	// 2.send message
	var sendCounter, successCounter, errorCount = 0, 0, 0
loop:
	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case producer.Input() <- &sarama.ProducerMessage{
			Topic: "async_topic",
			Value: sarama.StringEncoder("海绵宝宝 Go"),
		}: //发送消息
			sendCounter++
		case err := <-producer.Errors(): //处理错误
			log.Printf("Error producer: %s\n", err)
			errorCount++
		case suc := <-producer.Successes():
			log.Printf("Success to send, partition:%v , offset:%v\n", suc.Partition, suc.Offset)
			successCounter++
		case <-signals: //终止循环
			break loop
		}
	}
	// 统计结果
	log.Printf("Send counter: %v, error count: %v SuccessCounter: %v,\n", sendCounter, errorCount, successCounter)
}

// GoroutineAsyncSend 基于goroutine的异步
func GoroutineAsyncSend() {
	// 1.得到异步的producer
	broker := []string{"192.168.50.100:9092"}
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true //开启success channel来接受发送成功的信息
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
			log.Printf("Success to send, partition:%v , offset:%v\n", suc.Partition, suc.Offset)
			successCounter++
		}
	}()
	// 发送消息
	// 增加控制信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt) //ctrl + c
loop:
	for {
		time.Sleep(200 * time.Millisecond)
		select {
		case producer.Input() <- &sarama.ProducerMessage{
			Topic: "async_topic",
			Value: sarama.StringEncoder("海绵宝宝 Go"),
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
