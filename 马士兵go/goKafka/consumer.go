package goKafka

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func ConsumerTopic() {
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
	partitionConsumer, err := consumer.ConsumePartition("async_topic", 0, sarama.OffsetNewest) // 从最新的开始读消息
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
			log.Printf("Consumed Message: partition:%v offset:%v topic:%v value:%v \n", msg.Partition, msg.Offset, msg.Topic, string(msg.Value))
			consumerCounter++
		case <-signals:
			break loop
		}
	}
	log.Printf("Consumer Counter: %v\n", consumerCounter)
}
