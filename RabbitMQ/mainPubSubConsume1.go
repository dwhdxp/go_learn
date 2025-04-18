package main

import "learn/RabbitMQ/pubsub"

func main() {
	rabbitmq := pubsub.NewRabbitMQPubSub("newProduct")
	rabbitmq.ConsumePubSub()
}
