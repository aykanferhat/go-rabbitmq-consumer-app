package supplier

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Consumer struct {
}

func (c *Consumer) ConsumeSupplierUpdated(delivery amqp.Delivery) error {

	payload := string(delivery.Body)
	fmt.Printf("ConsumeSupplierUpdated, %s \n", payload)

	return nil
}
