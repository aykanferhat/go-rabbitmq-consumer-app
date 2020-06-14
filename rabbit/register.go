package rabbit

import (
	"errors"
	"github.com/streadway/amqp"
	"go-rabbitmq-consumer-app/config"
)

type QueueConsumerMap map[config.QueueConfig]func(delivery amqp.Delivery) error

var qcm QueueConsumerMap

func (c *Client) getRegisteredQueueConsumer() QueueConsumerMap {
	if qcm != nil {
		return qcm
	}
	queueConsumerMap := make(QueueConsumerMap)

	// Product Queue-Consumer binding
	queueConsumerMap[c.queuesConfig.Product.ProductCreated] = c.productConsumer.ConsumeProductCreated
	queueConsumerMap[c.queuesConfig.Product.ProductUpdated] = c.productConsumer.ConsumeProductUpdated
	queueConsumerMap[c.queuesConfig.Product.ProductStatusChanged] = c.productConsumer.ConsumeProductStatusChanged

	// Supplier Queue-Consumer binding
	queueConsumerMap[c.queuesConfig.Supplier.SupplierUpdated] = c.supplierConsumer.ConsumeSupplierUpdated

	qcm = queueConsumerMap
	return qcm
}

func FindConsumer(routingKey string) (func(delivery amqp.Delivery) error, error) {
	for key, value := range qcm {
		if key.RoutingKey == routingKey {
			return value, nil
		}
	}
	return nil, errors.New("Consumer not found, Routing Key: " + routingKey)
}
