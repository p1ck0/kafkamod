package kafkabroker

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type Broker struct {
	writer *kafka.Writer
	reader *kafka.Reader
}

func NewBroker(writer *kafka.Writer, reader *kafka.Reader) *Broker {
	return &Broker{
		writer: writer,
		reader: reader,
	}
}

func (b *Broker) Read(ctx context.Context, num int) {
	defer b.reader.Close()
	var mutex sync.RWMutex
	for {
		time.Sleep(1 * time.Second)
		mutex.RLock()
		m, err := b.reader.ReadMessage(context.Background())
		mutex.RUnlock()
		if err != nil {
			break
		}
		fmt.Printf("Reader №%d message at offset %d: %s = %s\n", num, m.Offset, string(m.Key), string(m.Value))
	}

	if err := b.reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func (b *Broker) Write(ctx context.Context, key []byte, value []byte) error {
	return b.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   key,
			Value: value,
		},
	)
}
