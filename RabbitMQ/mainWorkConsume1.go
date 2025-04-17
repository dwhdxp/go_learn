package main

import "learn/RabbitMQ/simple"

func main() {
	rabbitmq := simple.NewRabbitMQSimple("simple_queue")
	rabbitmq.ConsumeSimple()
}
