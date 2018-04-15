package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Region            string `envconfig:"AWS_REGION" default:"ap-southeast-2"`
	KinesisStreamName string `envconfig:"AWS_KINESIS_NAME" default:"test_endpoint"`
}

func NewConfig() *Config {
	c := &Config{}
	envconfig.MustProcess("", c)
	return c
}
