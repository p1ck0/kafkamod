package msgbroker

import (
	"context"
	"sync"

	"github.com/p1ck0/kafkamod/pkg/msgbroker/kafkabroker"
	"github.com/segmentio/kafka-go"
)

type MsgBroker interface {
	Write(ctx context.Context, key []byte, value []byte) error
	Read(ctx context.Context, num int, mutex *sync.RWMutex)
}

type Broker struct {
	MsgBroker MsgBroker
}

func NewMsgBroker(writer *kafka.Writer, reader *kafka.Reader) *Broker {
	return &Broker{
		MsgBroker: kafkabroker.NewBroker(writer, reader),
	}
}
