package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	logrus.Info("Starting RabbitMQ Metrics Exporter")
	SetupConfig()

	var c = new(Scheduler)

	err := c.Start()

	if err != nil {
		logrus.Error(err)
	}
}

func SetupConfig() {
	// Enable use of environment variables
	viper.SetEnvPrefix("RME")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".")

	viper.SetDefault("rabbitmq.host", "localhost")
	viper.SetDefault("rabbitmq.port", 15672)
	viper.SetDefault("rabbitmq.user", "guest")
	viper.SetDefault("rabbitmq.password", "guest")

	viper.SetDefault("exporter.type",  "statsd")
	viper.SetDefault("exporter.host",  "localhost")
	viper.SetDefault("exporter.port",  8125)


	viper.ReadInConfig()
}