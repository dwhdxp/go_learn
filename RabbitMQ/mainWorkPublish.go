package main

import (
	"fmt"
	"learn/RabbitMQ/simple"
	"strconv"
	"time"
)

func main() {
	// 工作模式只是在simple模式基础上进行了负载均衡
	rabbitmq := simple.NewRabbitMQSimple("simple_queue")
	for i := 0; i < 10; i++ {
		rabbitmq.PublishSimple("user word mode" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
	fmt.Println("publish success")
}
