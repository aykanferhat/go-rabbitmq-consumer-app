package config

type ApplicationConfig struct {
	Rabbit RabbitConfig `yaml:"rabbit"`
}

type RabbitConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	VirtualHost    string `yaml:"virtualHost"`
	ConnectionName string `yaml:"connectionName"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
}

type QueuesConfig struct {
	Product  ProductQueueConfig  `yaml:"product"`
	Supplier SupplierQueueConfig `yaml:"supplier"`
}

type ProductQueueConfig struct {
	ProductCreated       QueueConfig `yaml:"productCreated"`
	ProductUpdated       QueueConfig `yaml:"productUpdated"`
	ProductStatusChanged QueueConfig `yaml:"productStatusChanged"`
}

type SupplierQueueConfig struct {
	SupplierUpdated QueueConfig `yaml:"supplierUpdated"`
}

type QueueConfig struct {
	PrefetchCount int    `yaml:"prefetchCount"`
	ChannelCount  int    `yaml:"prefetchCount"`
	Exchange      string `yaml:"exchange"`
	ExchangeType  string `yaml:"exchangeType"`
	RoutingKey    string `yaml:"routingKey"`
	Queue         string `yaml:"queue"`
}
