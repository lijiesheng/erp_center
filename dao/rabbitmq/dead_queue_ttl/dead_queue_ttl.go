package dead_queue_ttl

import (
	dead_queue_ttl "erp_center/dao/rabbitmq/dead_queue_ttl/lib"
	"fmt"
)

var Rabbitmq *dead_queue_ttl.RabbitMQ

func Init(exchangeName string, routingKey string) {
	Rabbitmq = dead_queue_ttl.NewRabbitMQRouting(exchangeName, routingKey)
	fmt.Println("Rabbitmq==>", Rabbitmq)
}
