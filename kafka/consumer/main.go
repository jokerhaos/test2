package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

var address = []string{"myhao.com:29092", "myhao.com:29093", "myhao.com:29094"}

func main() {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	config.ClientID = "joker"
	config.Version = sarama.V3_3_1_0
	config.Consumer.Group.InstanceId = "group_mytest"
	consumer, err := sarama.NewConsumer(address, config)

	// sarama.NewConsumerGroup()

	if err != nil {
		fmt.Println("Failed to start consumer:", err)
		return
	}
	partitionList, err := consumer.Partitions("bigdata") // 通过topic获取到所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有的分区
		pc, err := consumer.ConsumePartition("bigdata", int32(partition), sarama.OffsetNewest) // 针对每个分区创建一个分区消费者
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) { // 为每个分区开一个go协程取值
			for msg := range pc.Messages() { // 阻塞直到有值发送过来，然后再继续等待
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
