package mq

type MessageBroker interface {
	SubscribeQueue(queueName string) error
	CreateQueue(queueName string) error
	SendMessage(queue string, message string) error
	ReceiveMessage(queue string) error
	DeleteMessage(queue string, messsageId string) error
}
