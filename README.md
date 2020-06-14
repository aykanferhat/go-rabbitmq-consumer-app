# go-rabbitmq-consumer-app


- First step
    ```
    docker run -d \
    --hostname rabbitmq \
    --name rabbitmq \
    -p 15672:15672 \
    -p 5672:5672 \
    -e RABBITMQ_DEFAULT_USER=user \
    -e RABBITMQ_DEFAULT_PASS=password \
    rabbitmq:3-management
    ```
        
- Second step
    ```
    go build
    ```
- Third step
    ```
    ./go-rabbitmq-consumer-app
    ```
