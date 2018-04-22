package services

import (
	"github.com/yunspace/serverless-golang/examples/aws-golang-dynamodb/config"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/satori/go.uuid"
	"github.com/yunspace/serverless-golang/examples/todo"
)

const TableName = "todos"

func givenConfig() *config.DynamoDBConfig {

	return &config.DynamoDBConfig{
		AWSConfig: config.AWSConfig{
			AWSRegion: "ap-southeast-2",
		},
		EndPoint: "http://localhost:4569",
		Table: TableName,
	}

}

func createTable(service *TodoDynamoDBService) error {
	//service := NewTodoDynamoDBService(givenConfig())
	_, err := service.db.Client().CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(TableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType: aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	return err
}

func dropTable(service *TodoDynamoDBService) error {
	_, err := service.db.Client().DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(TableName),
	})
	return err
}

func TestNewTodoDynamoDBService_canCreateAndListTable(t *testing.T) {
	// given
	dbConf := givenConfig()
	service := NewTodoDynamoDBService(dbConf)

	// when
	defer dropTable(service)
	assert.NoError(t, createTable(service))

	// then
	tables, err := service.db.ListTables().All()
	assert.NoError(t, err)
	assert.Equal(t, "Todos", tables[0])
}

func TestCreateAndGetTodo(t *testing.T) {
	// given
	id := uuid.NewV4()
	todo := &todo.Todo{
		Id: id,
	}
	dbConf := givenConfig()
	service := NewTodoDynamoDBService(dbConf)
	defer dropTable(service)
	assert.NoError(t, createTable(service))

	// when

	assert.NoError(t, service.Create(todo))

	// then
	created, err := service.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, id, created.ID())
}