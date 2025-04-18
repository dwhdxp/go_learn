package main

import "learn/RabbitMQ/routing"

func main() {
	rabbitmq2 := routing.NewRabbitMQRouting("routingEx", "rKeyTwo")
	rabbitmq2.ConsumeRouting()
}
