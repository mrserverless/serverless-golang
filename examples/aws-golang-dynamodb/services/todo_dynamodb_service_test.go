package services

import (
	"github.com/yunspace/serverless-golang/examples/aws-golang-dynamodb/config"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestNewTodoDynamoDBService_canConnect(t *testing.T) {
	// given
	dbConf := &config.DynamoDBConfig{
		AWSConfig: config.AWSConfig{
			AWSRegion: "ap-southeast-2",
		},
		EndPoint: "http://localhost:4569",
		Table: "todos",
	}

	// when
	service := NewTodoDynamoDBService(dbConf)
	_, err := service.db.ListTables().All()

	// then
	assert.NoError(t, err)
}