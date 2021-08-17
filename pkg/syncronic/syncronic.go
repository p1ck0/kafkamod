package syncronic

import (
	"context"
	"sync"

	"github.com/p1ck0/kafkamod/pkg/msgbroker"
)

type Syncronic struct {
	mutex  sync.RWMutex
	broker *msgbroker.Broker
}

func NewSyncronic(broker *msgbroker.Broker) *Syncronic {
	return &Syncronic{
		broker: broker,
	}
}

func (s *Syncronic) SyncWrite(ctx context.Context, key []byte, value []byte) error {
	s.mutex.Lock()
	err := s.broker.MsgBroker.Write(ctx, key, value)
	s.mutex.Unlock()
	return err
}

func (s *Syncronic) SyncRead(ctx context.Context, num int) {
	//time.Sleep(time.Second * 10)
	s.broker.MsgBroker.Read(ctx, num)
}
