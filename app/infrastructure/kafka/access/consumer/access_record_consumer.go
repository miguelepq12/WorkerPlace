package consumer

import (
	"WorkerPlace/app/application"
	"WorkerPlace/app/domain/entity"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Shopify/sarama"
)

type AccessRecordKafkaConsumer struct {
	worker sarama.Consumer
	usecase *application.AccessRecordUseCase
}

const (
	BrokersUrl         = "localhost:9092"
	AccessRecordTopics = "access-record"
)

func NewAccessRecordKafkaConsumer(usecase *application.AccessRecordUseCase) *AccessRecordKafkaConsumer {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// NewConsumer creates a new consumer using the given broker addresses and configuration
	conn, err := sarama.NewConsumer(strings.Split(BrokersUrl,","), config)
	if err != nil {
		return nil
	}
	
	return &AccessRecordKafkaConsumer{worker: conn,usecase: usecase}
}

func (c *AccessRecordKafkaConsumer)Run() {
	// calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := c.worker.ConsumePartition(AccessRecordTopics, 0,sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				CreateOrUpdate(msg,c.usecase)
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, msg.Topic, string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := c.worker.Close(); err != nil {
		panic(err)
	}
}

func CreateOrUpdate(msg *sarama.ConsumerMessage, usecase *application.AccessRecordUseCase){
	record := &entity.AccessRecord{}
	err := record.DecodeByte(msg.Value)
	if err != nil {
		fmt.Println("unable to decode message and create or update document")
	}else{
		usecase.CreateOrUpdateAccessRecord(record)
	}
}
