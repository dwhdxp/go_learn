package main

import (
	"fmt"
	"learn/RabbitMQ/routing"
	"strconv"
	"time"
)

func main() {
	rabbitmq1 := routing.NewRabbitMQRouting("routingEx", "rKeyOne")
	rabbitmq2 := routing.NewRabbitMQRouting("routingEx", "rKeyTwo")
	for i := 0; i < 10; i++ {
		rabbitmq1.PublishRouting("use routingEx, rKeyOne" + strconv.Itoa(i))
		rabbitmq2.PublishRouting("use routingEx, rKeyTwo" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	fmt.Println("publish done")
}
