# RabbitMQ metrics exporter

Tool to help export metrics from RabbitMQ to a metrics backend.

## Requirements

 - RabbitMQ with management plugin

## How to use

The easiest way to use this tool is to use the [docker image](https://hub.docker.com/r/maikelh/rabbitmq-metrics-exporter/).

An other options is to build the application yourself, see the building instructions below.

### Configuration

The tool use the excellent viper library which means you can supply the configuration in multiple formats. An example in
JSON is given below with all the possible options with default values.

```json
"rabbitmq" : {
    "host": "localhost",
    "port": 15672,
    "user": "guest",
    "password": "guest"
},
"exporter": {
    "type": "statsd",
    "host": "localhost",
    "port": 8125
}
```

The config should be named config[.ext] where ext is the type of your file.

It will try to find the config in the following locations:

- /etc/rabbitmq-metrics-exporter
- $HOME/.rabbitmq-metrics-exporter
- Executable directory

#### Environment values

It is also possible to use environment variables to set the config values. The enviroment variables need to have the `RME`
prefix.

So for example to set the RabbitMQ host we can set the following environment variable: `RME_RABBITMQ_HOST`


## How to build

TODO

