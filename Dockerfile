FROM alpine:3.5
ADD ["rabbitmq-metrics-exporter", "/app/rabbitmq-metrics-exporter"]
CMD ["/app/rabbitmq-metrics-exporter"]