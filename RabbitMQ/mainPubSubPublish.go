package main

import (
	"fmt"
	"learn/RabbitMQ/pubsub"
	"strconv"
	"time"
)

func main() {
	rabbitmq := pubsub.NewRabbitMQPubSub("newProduct")
	for i := 0; i < 10; i++ {
		rabbitmq.PublishPubSub("use pubsub mode" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	fmt.Println("publish done")
}
