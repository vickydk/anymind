package queue

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"anymind/pkg/shared/utils"
	"github.com/segmentio/kafka-go"
)

const (
	RetryConnectionMax   = 10
	ErrConnectionRefused = "connect: connection refused"
)

type consumerKafka struct {
	cfg         *ConsumerOptions
	readers     []*kafka.Reader
	kafkaConfig kafka.ReaderConfig
	receivers   map[string]Receiver
	ctx         context.Context
}

func SetupKafkaConsumer(cfg *ConsumerOptions) *consumerKafka {
	fmt.Println("Try NewKafkaConsumer ...")
	return &consumerKafka{
		cfg:       cfg,
		receivers: map[string]Receiver{},
		ctx:       context.Background(),
	}
}

func (c *consumerKafka) Start() *consumerKafka {
	brokerList := strings.Split(c.cfg.Brokers, ",")

	c.kafkaConfig = kafka.ReaderConfig{
		Brokers:                brokerList,
		GroupID:                c.cfg.Group,
		MinBytes:               1,
		MaxBytes:               10e6,                   // 10MB
		MaxWait:                100 * time.Millisecond, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval:        -1,
		StartOffset:            kafka.LastOffset,
		WatchPartitionChanges:  true,
		PartitionWatchInterval: time.Minute,
	}

	c.setupReader(c.cfg.Topics)

	if c.readers == nil || len(c.readers) == 0 {
		panic("kafka consumer group reader is nil")
	}

	return c
}

func (c *consumerKafka) SetReceivers(receivers map[string]Receiver) {
	c.receivers = receivers
}

func (c *consumerKafka) SetReceiver(topic string, receiver Receiver) {
	c.receivers[topic] = receiver
}

func (c *consumerKafka) SetContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *consumerKafka) setupReader(topics []ConsumerTopic) *consumerKafka {
	c.readers = make([]*kafka.Reader, 0)
	for _, topic := range topics {
		if topic.Topic != "" {
			copyKafkaConfig := c.kafkaConfig
			copyKafkaConfig.Topic = topic.Topic

			reader := kafka.NewReader(copyKafkaConfig)

			go c.startReader(reader)

			c.readers = append(c.readers, reader)
		}
	}
	return c
}

func (c *consumerKafka) startReader(reader *kafka.Reader) {
	retry := 0
	for {
		msg, err := reader.ReadMessage(c.ctx)
		if err != nil {
			fmt.Printf("consumer reading message : %s\n", err.Error())
			if err == context.Canceled || err == context.DeadlineExceeded {
				break
			}

			if strings.Contains(err.Error(), ErrConnectionRefused) || err == io.EOF {
				if retry >= RetryConnectionMax {
					break
				}
				retry++
				time.Sleep(utils.Default.Duration(retry))
			}
			continue
		}
		retry = 0

		// broadcast
		receiver, ok := c.receivers[msg.Topic]
		if !ok {
			// ignore this topic
			continue
		}

		go receiver(string(msg.Key), msg.Topic, msg.Value)

	}
}

func (c *consumerKafka) closeReader() *consumerKafka {
	for _, reader := range c.readers {
		_ = reader.Close()
	}
	return c
}

func (c *consumerKafka) Stop() error {
	errFound := false
	var errs []string

	for i := range c.readers {
		if err := c.readers[i].Close(); err != nil {
			errFound = true
			errs = append(errs, err.Error())
		}
	}

	if errFound {
		return fmt.Errorf("closing consumer readers : %+v", strings.Join(errs, " | "))
	}

	return nil
}
