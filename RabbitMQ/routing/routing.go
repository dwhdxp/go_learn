package routing

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// url地址：amqp://用户名:密码@IP地址:端口号/vhost
const MQURL = "amqp://xp:123456@127.0.0.1:5672/dxp"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

// failOnError 错误处理函数
func (r *RabbitMQ) failOnError(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// Close 断开连接和通道
func (r *RabbitMQ) Close() {
	r.conn.Close()
	r.channel.Close()
}

// NewRabbitMQ 初始化RabbitMQ
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
	var err error
	// 获取Connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnError(err, "Failed to connect to RabbitMQ")
	// 获取Channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnError(err, "Failed to open a channel")
	return rabbitmq
}

/*
	Routing：Direct类型交换器，根据路由键匹配规则将消息投递到对应的队列中
	在发布订阅模式基础上，更改交换器类型，并指定routingKey
*/
// NewRabbitMQRouting 初始化路由模式的RabbitMQ实例
func NewRabbitMQRouting(exchangeName, routingKey string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, routingKey)
}

// PublishRouting 路由模式生产者
func (r *RabbitMQ) PublishRouting(message string) {
	// 1.尝试创建Direct类型交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		// 若为true，则交换机无法被client用来推送消息，只能进行exchange之间的绑定
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.failOnError(err, "PublishRouting failed to declare a exchange")
	}

	// 2.发布消息到队列
	r.channel.Publish(
		r.Exchange,
		r.Key,
		// 如果为true，当exchange找不到符合要求的队列时，会将消息返回给生产者
		false,
		// 如果为true，当exchange发送消息到队列后发现队列上没有消费者，会将消息返回给生产者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// ConsumeRouting 路由模式消费者
func (r *RabbitMQ) ConsumeRouting() {
	// 1.尝试创建Direct类型交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		// 若为true，则交换机无法被client用来推送消息，只能进行exchange之间的绑定
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.failOnError(err, "PublishPubSub failed to declare a exchange")
	}

	// 2.尝试创建队列，并绑定到交换机
	q, err := r.channel.QueueDeclare(
		"", // 随机生成队列
		// 是否持久化
		false,
		// 当最后一个消费者退订后是否自动删除
		false,
		// 是否具有排他性
		true,
		// 是否阻塞
		false,
		// 额外的参数
		nil,
	)
	if err != nil {
		r.failOnError(err, "ConsumePubSub failed to declare a queue")
	}

	// 绑定到exchange
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	// 2.接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",
		// 是否自动确认应答，若为true，则消息一旦被消费，RabbitMQ会自动从队列中删除消息
		true,
		// 是否独占，若为true，则该消费者为队列的唯一消费者
		false,
		// 若为true，则不能将生产者发送的消息传递给同一个Connection中的消费者
		false,
		// 是否阻塞，若为true，则当队列中没有消息时，消费者会阻塞等待
		false,
		// 额外的参数
		nil,
	)
	if err != nil {
		r.failOnError(err, "ConsumePubSub failed to register a consumer")
	}

	// 3.启用协程处理消息
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			// 处理消息逻辑
			log.Printf("Received a message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
