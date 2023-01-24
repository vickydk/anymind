package config

import (
	"fmt"

	Database "anymind/pkg/shared/database"
	Queue "anymind/pkg/shared/queue"

	"github.com/spf13/viper"
)

type Config struct {
	Apps                 Apps                  `json:"apps"`
	Database             DB                    `json:"database"`
	Kafka                *KafkaConfig          `json:"kafka"`
	KafkaProcedureTopics *KafkaProcedureTopics `json:"kafkaProcedureTopics"`
	KafkaConsumerTopics  *KafkaConsumerTopics  `json:"kafkaConsumerTopics"`
}

type Apps struct {
	Name     string `json:"name"`
	HttpPort int    `json:"httpPort"`
	GRPCPort int    `json:"grpcPort"`
	Version  string `json:"version"`
}

type DB struct {
	Master Database.ConfigDatabase `json:"master"`
	Slave  Database.ConfigDatabase `json:"slave"`
}

type QueueTopic struct {
	Name  string `json:"name"`
	Topic string `json:"topic"`
}

type KafkaConfig struct {
	Consumer struct {
		Brokers string       `json:"brokers"`
		Group   string       `json:"group"`
		Topics  []QueueTopic `json:"topics"`
	} `json:"consumer"`
	Producer Queue.ProducerOptions `json:"producer"`
}

type KafkaProcedureTopics struct {
	HistoryTransaction string `json:"historyTransaction"`
}

type KafkaConsumerTopics struct {
	HistoryTransaction string `json:"historyTransaction"`
}

func (c *Config) AppAddress() string {
	return fmt.Sprintf(":%v", c.Apps.HttpPort)
}

func NewConfig(path string) *Config {
	fmt.Println("Try NewConfig ... ")

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf := Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
