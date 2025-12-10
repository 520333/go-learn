package goKafka

import (
	"log"

	"github.com/IBM/sarama"
)

func SyncSend() {
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
	// 2.设置消息内容
	msg := &sarama.ProducerMessage{
		Topic: "sync_topic",
		Value: sarama.StringEncoder("Dawn Go Kafka"),
		//Value: sarama.ByteEncoder([]byte("hello world")),
	}
	// 3.发送消息 返回: 分区索引,偏移量,错误
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Send Failed err:%s\n", err)
	} else {
		log.Printf("Send Success partition:%d offset:%d\n", partition, offset)
	}
}
