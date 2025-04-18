package main

import "learn/RabbitMQ/routing"

func main() {
	rabbitmq1 := routing.NewRabbitMQRouting("routingEx", "rKeyOne")
	rabbitmq1.ConsumeRouting()
}
