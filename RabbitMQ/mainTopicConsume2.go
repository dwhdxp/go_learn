package main

import "learn/RabbitMQ/topic"

func main() {
	rabbitmq := topic.NewRabbitMQTopic("topic_Ex", "rKey.*.two")
	rabbitmq.ConsumeTopic()
}
