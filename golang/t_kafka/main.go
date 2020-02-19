package main

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/qingsong-he/ce"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func init() {
	ce.Print(os.Args[0])
}

var kafkaAddr = []string{"localhost:9092"}

func kafkaProduceMsgBySarama(topicName string) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	err := cfg.Validate()
	ce.CheckError(err)

	var wg sync.WaitGroup

	producer, err := sarama.NewAsyncProducer(kafkaAddr, cfg)
	ce.CheckError(err)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range producer.Successes() {
			ce.Print(msg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			ce.Print(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		producer.Input() <- &sarama.ProducerMessage{
			Topic: topicName,
			Value: sarama.StringEncoder(time.Now().Format(time.RFC3339)),
		}
	}()

	var mainByExitAlarm chan os.Signal
	mainByExitAlarm = make(chan os.Signal, 1)
	signal.Notify(mainByExitAlarm, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)

forLableByNotify:
	for {
		s := <-mainByExitAlarm
		ce.Info(s.String())
		switch s {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			producer.AsyncClose()
			break forLableByNotify

		case syscall.SIGHUP:
		default:
			producer.AsyncClose()
			break forLableByNotify
		}
	}

	wg.Wait()
}

type exampleConsumerGroupHandler struct{}

func (*exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (*exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (*exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		ce.Print("consume:", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func kafkaConsumeMsgBySarama(groupIdByConsume string, topicNames ...string) {
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Version = sarama.V2_4_0_0
	err := cfg.Validate()
	ce.CheckError(err)

	var wg sync.WaitGroup

	consumerByGroup, err := sarama.NewConsumerGroup(kafkaAddr, groupIdByConsume, cfg)
	ce.CheckError(err)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range consumerByGroup.Errors() {
			ce.Print(err)
		}
	}()

	wg.Add(1)
	handler := &exampleConsumerGroupHandler{}
	go func() {
		defer func() {
			recover()
		}()
		defer wg.Done()
		for {
			err := consumerByGroup.Consume(context.Background(), topicNames, handler)
			ce.Print("consume +1")
			ce.CheckError(err)
		}
	}()

	var mainByExitAlarm chan os.Signal
	mainByExitAlarm = make(chan os.Signal, 1)
	signal.Notify(mainByExitAlarm, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)

forLableByNotify:
	for {
		s := <-mainByExitAlarm
		ce.Info(s.String())
		switch s {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			consumerByGroup.Close()
			break forLableByNotify

		case syscall.SIGHUP:
		default:
			consumerByGroup.Close()
			break forLableByNotify
		}
	}

	wg.Wait()
}

func main() {
	if len(os.Args) >= 3 {
		switch os.Args[1] {
		case "p":
			kafkaProduceMsgBySarama(os.Args[2])
		case "c":
			kafkaConsumeMsgBySarama(os.Args[2], os.Args[3:]...)
		}
	}
}
