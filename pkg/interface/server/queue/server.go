package queue

import (
	"anymind/pkg/interface/container"
	Queue "anymind/pkg/shared/queue"
)

func StartSchedulerService(container *container.Container) {
	handlers := SetupHandlers(container)

	topics := make([]Queue.ConsumerTopic, 0)
	for _, t := range container.Config.Kafka.Consumer.Topics {
		topics = append(topics, Queue.ConsumerTopic{
			Name:  t.Name,
			Topic: t.Topic,
		})
	}

	queueHandler := Queue.SetupKafkaConsumer(&Queue.ConsumerOptions{
		Brokers: container.Config.Kafka.Consumer.Brokers,
		Group:   container.Config.Kafka.Consumer.Group,
		Topics:  topics,
	})
	// notification
	queueHandler.SetReceiver(container.Config.KafkaConsumerTopics.HistoryTransaction, handlers.history.HistoryTransaction)

	queueHandler.Start()
}
