package goKafka

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func CreateTopic() {
	var addr = "192.168.50.100:9092"
	broker := sarama.NewBroker(addr)
	// 打开Broker 建立连接
	config := sarama.NewConfig()
	if err := broker.Open(config); err != nil {
		log.Fatalln(err)
	}
	defer func(broker *sarama.Broker) {
		_ = broker.Close()
	}(broker)
	// 1.2判定连接状态 open会连接但是非阻塞连接模式 不会等到连接成功在返回 通常会强制连接  broker.Connected()
	connected, err := broker.Connected()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Connected to Kafka broker %v \n", connected)

	// 2.设置主题
	topicKey := "topic_more_partition_1"
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     3, //分区数了
		ReplicationFactor: 1, //复制因子
	}

	// 3.发出创建主题的请求 请求对象
	request := sarama.CreateTopicsRequest{
		TopicDetails: map[string]*sarama.TopicDetail{
			topicKey: topicDetail,
		},
		Timeout:      time.Second * 1,
		ValidateOnly: false,
	}
	// 发出请求
	response, err := broker.CreateTopics(&request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("CreateTopic response : %v \n", response)

	// 4.查看分区数量
	consumer, err := sarama.NewConsumer([]string{addr}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(consumer sarama.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(consumer)
	partitions, err := consumer.Partitions(topicKey)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Partitions: %v \n", partitions)
}
