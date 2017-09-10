package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Environment string `envconfig:"ENV"`
	DB          *DynamoDBConfig
}

type AWSConfig struct {
	AWSRegion string `envconfig:"AWS_REGION"`
}

type DynamoDBConfig struct {
	AWSConfig
	Table    string `envconfig:"DYNAMODB_TABLE"`
	EndPoint string `envconfig:"DYNAMODB_ENDPOINT"`
}

func NewConfig() *Config {
	c := &Config{}
	envconfig.MustProcess("", c)
	return c
}
