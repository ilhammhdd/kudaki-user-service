package usecases

import (
	"os"

	sarama "gopkg.in/Shopify/sarama.v1"
)

type EventSourceProducer interface {
	Set(topic string, partition int32, offset int64)
	Get() (topic string, partition int32, offset int64)
	SyncProduce(key string, value []byte) (producedPartition int32, producedOffset int64, err error)
}

type EventSourceConsumer interface {
	Set(topic string, partition int32, offset int64)
	Get() (topic string, partition int32, offset int64)
	Consume() (partCons sarama.PartitionConsumer, signals chan os.Signal, close chan bool)
}
