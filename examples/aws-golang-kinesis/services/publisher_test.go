package services

import (
	"testing"
	"github.com/yunspace/serverless-golang/examples/aws-golang-kinesis/config"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestDataPublisher(t *testing.T) {
	// given
	svc := NewKinesisService(config.NewConfig())
	bytes, _ := ioutil.ReadFile("../event.json")


	// when
	result, err := svc.Publish(bytes)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

