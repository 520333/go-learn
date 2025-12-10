package goKafka

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
)

// Event 需要发送的消息类型
type Event struct {
	Name   string    `json:"name"`
	Type   string    `json:"type"`
	Source string    `json:"source"`
	Target string    `json:"target"`
	Time   time.Time `json:"time"`
}
type EventEncoder Event

func (e EventEncoder) Encode() ([]byte, error) {
	// 将Event类型数据做JSON编码
	j, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (e EventEncoder) Length() int {
	j, err := json.Marshal(e)
	if err != nil {
		return 0
	}
	return len(j)
}

func EventProduce() {
	// 1.获取同步生产者
	broker := []string{"192.168.50.100:9092"}
	producer, err := sarama.NewSyncProducer(broker, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// 模拟数据
	evt := Event{
		Name:   "user",
		Type:   "buy",
		Source: "42",
		Target: "10098",
		Time:   time.Now(),
	}
	// 2.设置消息内容
	msg := &sarama.ProducerMessage{
		Topic: "event_topic",
		Value: EventEncoder(evt),
	}
	// 3.发送消息 返回: 分区索引,偏移量,错误
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Send Failed err:%s\n", err)
	} else {
		log.Printf("Send Success partition:%d offset:%d\n", partition, offset)
	}
}

func EventConsumerTopic() {
	// 1.获取消费者
	consumer, err := sarama.NewConsumer([]string{"192.168.50.100:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(consumer sarama.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(consumer)

	// 2.获取topic 具体某个partition的消费者
	// topic 1:N partition
	partitionConsumer, err := consumer.ConsumePartition("event_topic", 0, sarama.OffsetNewest) // 从最新的开始读消息
	if err != nil {
		log.Fatalln(err)
	}
	defer func(pc sarama.PartitionConsumer) {
		if err := pc.Close(); err != nil {
			log.Fatalln(err)
		}
	}(partitionConsumer)

	// 3.消费操作
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	var consumerCounter int
loop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf(
				"Consumed Message: partition:%v offset:%v topic:%v value:%v type:%T \n",
				msg.Partition, msg.Offset, msg.Topic, string(msg.Value), msg.Value)
			evt := Event{}
			err := json.Unmarshal(msg.Value, &evt)
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("数据:%v 类型:%T\n", evt, evt)
			consumerCounter++
		case <-signals:
			break loop
		}
	}
	log.Printf("Consumer Counter: %v\n", consumerCounter)
}
