package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/p1ck0/kafkamod/internal/kafkabroker"
	"github.com/p1ck0/kafkamod/pkg/msgbroker"
	"github.com/p1ck0/kafkamod/pkg/syncronic"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	writer := kafkabroker.NewKafkaWriter()
	reader := kafkabroker.NewKafkaReader()
	broker := msgbroker.NewMsgBroker(writer, reader)
	sr := syncronic.NewSyncronic(broker)
	var wg sync.WaitGroup
	wg.Add(10)
	go syncW(sr, 6, wg)
	go syncW(sr, 7, wg)
	go syncW(sr, 8, wg)
	go syncW(sr, 9, wg)
	go syncW(sr, 10, wg)
	go syncR(sr, 1, wg)
	go syncR(sr, 2, wg)
	go syncR(sr, 3, wg)
	go syncR(sr, 4, wg)
	go syncR(sr, 5, wg)
	wg.Wait()
}

func syncR(sr *syncronic.Syncronic, num int, wg sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()
	sr.SyncRead(ctx, num)
}

func syncW(sr *syncronic.Syncronic, num int, wg sync.WaitGroup) {
	defer wg.Done()
	ctx := context.Background()
	for {
		time.Sleep(time.Second * 1)
		err := sr.SyncWrite(ctx, []byte(fmt.Sprintf("key %d", rand.Uint64())), []byte(fmt.Sprintf("it is gorotine number %d", num)))
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Writer №%d\n", num)
	}
}
