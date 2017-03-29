FROM golang:1.8.0
ADD ["rabbitmq-metrics-exporter", "/app/rabbitmq"]
CMD ["/app/rabbitmq"]