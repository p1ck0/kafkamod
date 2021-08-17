package kafkabroker

import (
	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP("localhost:29092"),
		Topic:    "n511test-topic",
		Balancer: &kafka.LeastBytes{},
		//Async:    true,
	}
}

func NewKafkaReader() *kafka.Reader {
	batchSize := int(10e8)
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092"},
		Topic:     "n511test-topic",
		Partition: 0,
		MinBytes:  batchSize,
		MaxBytes:  batchSize,
		//ReadBackoffMax: time.Duration(time.Second * 5),
	})
}
