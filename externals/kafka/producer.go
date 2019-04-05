package kafka

import (
	"os"
	"strings"
	"time"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type Production struct {
	topic     string
	partition int32
	offset    int64
}

func NewProduction() *Production { return &Production{} }

func (p *Production) Set(topic string, partition int32, offset int64) {
	p.topic = topic
	p.partition = partition
	p.offset = offset
}

func (p *Production) Get() (string, int32, int64) {
	return p.topic, p.partition, p.offset
}

func (p *Production) SyncProduce(key string, value []byte) (partition int32, offset int64, err error) {
	type rtrn struct {
		Successes bool
		Errors    bool
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewManualPartitioner
	config.Producer.Return = rtrn{
		Successes: true,
		Errors:    true}

	prod, err := sarama.NewSyncProducer(strings.Split(os.Getenv("KAFKA_BROKERS"), ","), config)
	errorkit.ErrorHandled(err)

	defer prod.Close()

	t, part, o := p.Get()

	msg := sarama.ProducerMessage{
		Topic:     t,
		Offset:    o,
		Partition: part,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(value),
		Timestamp: time.Now(),
	}

	return prod.SendMessage(&msg)
}
