package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	KinesisStreamName string `envconfig:"AWS_KINESIS_NAME"`
}

func NewConfig() *Config {
	c := &Config{}
	envconfig.MustProcess("", c)
	return c
}
