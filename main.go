package main

import (
	"fmt"
	"go-rabbitmq-consumer-app/config"
	"go-rabbitmq-consumer-app/product"
	"go-rabbitmq-consumer-app/rabbit"
	"go-rabbitmq-consumer-app/supplier"
)

func main() {

	configurationManager := config.NewConfigurationManager()
	rabbitConfig := configurationManager.GetRabbitConfig()
	queuesConfig := configurationManager.GetQueuesConfig()

	productConsumer := product.Consumer{}
	supplierConsumer := supplier.Consumer{}

	rabbitClient := rabbit.NewRabbitClient(rabbitConfig, queuesConfig, productConsumer, supplierConsumer)
	defer rabbitClient.CloseConnection()

	rabbitClient.DeclareExchangeQueueBindings()

	consumerChan := make(chan bool)

	rabbitClient.InitializeConsumers()
	fmt.Println("Started consumers")

	<-consumerChan // close(consumerChan)
}
