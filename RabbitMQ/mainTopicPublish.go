package main

import (
	"fmt"
	"learn/RabbitMQ/topic"
	"strconv"
	"time"
)

func main() {
	rabbitmq1 := topic.NewRabbitMQTopic("topic_Ex", "rKey.topic.one")
	rabbitmq2 := topic.NewRabbitMQTopic("topic_Ex", "rKey.topic.two")
	for i := 0; i < 10; i++ {
		rabbitmq1.PublishTopic("use topic mode one" + strconv.Itoa(i))
		rabbitmq2.PublishTopic("use topic mode two" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	fmt.Println("publish done")
}
