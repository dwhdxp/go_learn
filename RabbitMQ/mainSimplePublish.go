package main

import (
	"fmt"
	"learn/RabbitMQ/simple"
)

func main() {
	rabbitmq := simple.NewRabbitMQSimple("simple_queue")
	rabbitmq.PublishSimple("use simple mode")
	fmt.Println("publish message ok")
}
