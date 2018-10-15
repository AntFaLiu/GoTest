package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/streadway/amqp"
)

/**
	AMQP协议预定义了14个属性，大多数我们都很少用到，以下几个是比较常用的。
	persistent：消息持久性
	content_type：用来描述编码的MIME类型
	reply_to：回调queue的名字
	correlation_id：将远程RPC请求，进行关联的唯一标识

correlation_id
  如果为每个RPC的请求创建一个queue效率是非常低的，正常发送到queue的一个Message，它不知道是从哪里发过来的，
  而correlation_id属性的存在就是为每个请求设置一个唯一值，在回调接收消息的时候，也会带回这个属性进行匹配，如果不匹配，这个消息就不会被处理。

	接下来我们将使用RabbitMQ搭建一个RPC系统：一个客户端和一个可扩展的RPC服务器，RPC的工作流程如下：

	客户端启动时，创建一个匿名的exclusive callback queue
	客户端发送请求时，要带两个属性reply_to（设置回调的queue）和correlation_id(唯一标识)
	将请求发送到一个RPC queue
	RPC的server端 ，一直在等待请求，当消息到达时会对过reply_to回复到指定的queue
	客户端在等queue从server的回调，检查 correlation_id是否一致，如果和请求时发送的一致，则做其他响应
 */
const (
	//AMQP URI
	uri = "amqp://guest:guest@localhost:5672/"
	//Durable AMQP queue name
	queueName = "rpc-queue"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	//调用发布消息函数
	publish(uri, queueName)
}

//发布者的方法
//
//@amqpURI, amqp的地址
//@queue, queue的名称
func publish(amqpURI string, queue string) {
	//建立连接
	log.Printf("dialing %q", amqpURI)
	connection, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer connection.Close()

	//创建一个Channel
	log.Printf("got Connection, getting Channel")
	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	//创建一个queue
	log.Printf("got queue, declaring %q", queue)
	q, err := channel.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//均衡处理，每次处理一条消息
	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	//订阅一个消息
	//log.Printf("Queue bound to Exchange, starting Consume")
	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	// 发布消息
	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] server端接收到的数据是 (%d)", n)
			response := n * 2

			err = channel.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")

	//没有写入数据，一直等待读，阻塞当前线程，目的是让线程不退出
	<-forever
}
