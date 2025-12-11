package goKafka

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

func ConsumerPartition() {
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

	// 2.获取topic下的分区列表
	topic := "topic_more_partition_1"
	partitions, err := consumer.Partitions(topic)
	log.Println(partitions)
	if err != nil {
		log.Fatalln(err)
	}
	// 3.遍历每个分区
	var consumerCounter int
	for _, partition := range partitions {
		// 获取每个分区的消费者
		partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest) // 从最新的开始读消息
		if err != nil {
			log.Fatalln(err)
		}
		//4.利用goroutine完成消费
		go func(pc sarama.PartitionConsumer) {
			// 关闭
			defer func() {
				if err := pc.Close(); err != nil {
					log.Fatalln(err)
				}
			}()
			// 消费
			for msg := range pc.Messages() {
				log.Println(pc)
				log.Printf("Consumed Message: partition:%v offset:%v topic:%v value:%v \n", msg.Partition, msg.Offset, msg.Topic, string(msg.Value))
				consumerCounter++
			}
		}(partitionConsumer)

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		select {
		case <-signals:
		}
		log.Printf("Consumer Counter: %v\n", consumerCounter)
	}
}
