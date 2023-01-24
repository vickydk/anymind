package queue

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/segmentio/encoding/json"
)

type producerKafka struct {
	producer sarama.SyncProducer
}

func SetupKafkaProducer(options *ProducerOptions) Producer {
	fmt.Println("Try NewKafkaProducer ...")

	list := strings.Split(options.Address, ",")
	cfg := sarama.NewConfig()
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = options.RetryMax
	cfg.Producer.Return.Successes = options.ReturnSuccesses

	p, err := sarama.NewSyncProducer(list, cfg)
	if err != nil {
		panic("Could not create sarama producer: " + err.Error())
	}

	return &producerKafka{
		producer: p,
	}
}

func (p *producerKafka) setProducer(producer sarama.SyncProducer) {
	p.producer = producer
}

func (p *producerKafka) SendMessage(key, topic string, message interface{}) (err error) {
	var data []byte
	if s, ok := message.(string); ok {
		data = []byte(s)
	} else if b, ok := message.([]byte); ok {
		data = b
	} else {
		data, err = json.Marshal(message)
		if err != nil {
			err = errors.New("error put to queue")
			return err
		}
	}

	msg := &sarama.ProducerMessage{
		Key:       sarama.StringEncoder(key),
		Topic:     topic,
		Partition: -1,
		Timestamp: time.Now(),
		Value:     sarama.ByteEncoder(data),
	}

	_, _, err = p.producer.SendMessage(msg)
	if err != nil {
		err = errors.New("error put to queue")
		return err
	}

	return
}
