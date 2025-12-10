package goKafka

import (
	"log"

	"github.com/IBM/sarama"
)

func Ack() {
	broker := []string{"192.168.50.100:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal //1 default
	config.Producer.RequiredAcks = sarama.WaitForAll   //-1
	config.Producer.RequiredAcks = sarama.NoResponse   //0
	producer, err := sarama.NewSyncProducer(broker, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
}
