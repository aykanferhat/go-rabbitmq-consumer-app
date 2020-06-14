package product

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
}

func (c *Consumer) ConsumeProductCreated(delivery amqp.Delivery) error {

	payload := string(delivery.Body)
	fmt.Printf("ConsumeProductCreated, %s \n", payload)

	return nil
}

func (c *Consumer) ConsumeProductUpdated(delivery amqp.Delivery) error {

	payload := string(delivery.Body)
	fmt.Printf("ConsumeProductUpdated, %s \n", payload)

	return nil
}

func (c *Consumer) ConsumeProductStatusChanged(delivery amqp.Delivery) error {

	payload := string(delivery.Body)
	fmt.Printf("ConsumeProductStatusChanged, %s \n", payload)

	return nil
}
