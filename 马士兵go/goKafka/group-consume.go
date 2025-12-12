package goKafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

func GroupConsume() {
	// 1.创建消费者组
	addrs := []string{"192.168.50.100:9092"}
	groupID := "DawnGroup"
	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
	// 消费信道返回
	group, err := sarama.NewConsumerGroup(addrs, groupID, conf)

	// 设置分配策略
	conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		//sarama.NewBalanceStrategyRange(),//默认
		//sarama.NewBalanceStrategyRoundRobin(), //轮询
		sarama.BalanceStrategySticky, //粘性
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		_ = group.Close()
	}()

	// 2.处理group的错误
	go func() {
		for err := range group.Errors() {
			log.Println(err)
		}
	}()

	// 3.组消费
	// 带有cancel context
	ctx, cancel := context.WithCancel(context.Background())
	topics := []string{"topic_more_partition_1"}
	wg := &sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			// 组内消费者成员改变时，重新执行，重新分配
			handler := GroupConsumerHandler{}
			if err := group.Consume(ctx, topics, handler); err != nil {
				log.Println(err)
			}
			// 判定context 是否cancel
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
		}
	}()
	// 信号阻塞
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
		// 终止
		cancel()
	}
	wg.Wait()
}

// GroupConsumerHandler 定义组消费处理器 group consume handler
type GroupConsumerHandler struct {
}

// Setup 重新消费时执行 增减组内消费者时
func (GroupConsumerHandler) Setup(cgs sarama.ConsumerGroupSession) error {
	log.Println("Setup")
	log.Println(cgs.Claims()) //该消费者所分配的消费分区
	cgs.ResetOffset("topic_more_partition_1", 0, 2048, "")
	return nil
}

// Cleanup 当组内消费者推出时执行
func (GroupConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("Cleanup")
	return nil
}

// ConsumeClaim 组消费的核心方法
func (GroupConsumerHandler) ConsumeClaim(cgs sarama.ConsumerGroupSession, cgc sarama.ConsumerGroupClaim) error {
	log.Println("ConsumeClaim")
	// 消费
	for msg := range cgc.Messages() {
		log.Printf("Consumed Message: partition:%v offset:%v topic:%v \n", msg.Partition, msg.Offset, msg.Topic)
		// 标记该消息已经被消费
		cgs.MarkMessage(msg, "")
	}
	return nil
}
