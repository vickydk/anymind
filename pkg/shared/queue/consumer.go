package queue

type Receiver func(threadID, topic string, msg []byte)

type Consumer interface {
	SetReceivers(receivers map[string]Receiver)
	SetReceiver(topic string, receiver Receiver)
}
