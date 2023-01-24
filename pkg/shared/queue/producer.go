package queue

type Producer interface {
	SendMessage(key, topic string, message interface{}) (err error)
}
