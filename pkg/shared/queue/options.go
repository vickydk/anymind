package queue

type ProducerOptions struct {
	Address         string `json:"address"`
	RetryMax        int    `json:"retryMax"`
	ReturnSuccesses bool   `json:"returnSuccesses"`
}

type ConsumerOptions struct {
	Brokers string          `json:"brokers"`
	Group   string          `json:"group"`
	Topics  []ConsumerTopic `json:"topics"`
}

type ConsumerTopic struct {
	Name  string `json:"name"`
	Topic string `json:"topic"`
}
