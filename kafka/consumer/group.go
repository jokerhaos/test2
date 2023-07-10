package main

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
)

var address = []string{"myhao.com:29092", "myhao.com:29093", "myhao.com:29094"}

// // 实现  github.com/Shopify/sarama/consumer_group.go/ConsumerGroupHandler  这个接口
type AAAConsumerGroupHandler struct {
}

func (AAAConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (AAAConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// 这个方法用来消费消息的
func (h AAAConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// 获取消息
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		fmt.Println("msg key: ", string(msg.Key))
		fmt.Println("msg value: ", string(msg.Value))
		// 将消息标记为已使用
		sess.MarkMessage(msg, "")
	}
	return nil
}

// 接收数据
func main() {
	// 先初始化 kafka
	config := sarama.NewConfig()
	// Version 必须大于等于  V0_10_2_0
	config.Version = sarama.V3_3_1_0
	config.Consumer.Return.Errors = true
	fmt.Println("start connect kafka")
	// 开始连接kafka服务器
	group, err := sarama.NewConsumerGroup(address, "group_mytest2", config)

	if err != nil {
		fmt.Println("connect kafka failed; err", err)
		return
	}
	// 检查错误
	go func() {
		for err := range group.Errors() {
			fmt.Println("group errors : ", err)
		}
	}()

	ctx := context.Background()
	fmt.Println("start get msg")
	// for 是应对 consumer rebalance
	for {
		// 需要监听的主题
		topics := []string{"bigdata"}
		handler := AAAConsumerGroupHandler{}
		// 启动kafka消费组模式，消费的逻辑在上面的 ConsumeClaim 这个方法里
		err := group.Consume(ctx, topics, handler)

		if err != nil {
			fmt.Println("consume failed; err : ", err)
			return
		}
	}

}
