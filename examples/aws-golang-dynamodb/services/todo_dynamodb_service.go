package services

import (
	"github.com/guregu/dynamo"
	"github.com/yunspace/serverless-golang/api"
	"github.com/yunspace/serverless-golang/examples/todo"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"github.com/yunspace/serverless-golang/examples/aws-golang-dynamodb/config"
)

var _ api.CRUDAPI = (*TodoDynamoDBService)(nil)

const (
	partitionKey = "ID"
)

type TodoDynamoDBService struct {
	db        *dynamo.DB
	table     *dynamo.Table
}

func NewDynamoDB(dbConf *config.DynamoDBConfig) *dynamo.DB {
	s, err := session.NewSession()
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	awsConfig := aws.NewConfig()
	awsConfig.Region = aws.String(dbConf.AWSRegion)
	if dbConf.EndPoint != "" {
		awsConfig.Endpoint = aws.String(dbConf.EndPoint)
	}
	return dynamo.New(s, awsConfig)
}

func NewTodoDynamoDBService(dbConf *config.DynamoDBConfig) *TodoDynamoDBService {
	db := NewDynamoDB(dbConf)
	todoTable := db.Table(dbConf.Table)
	return &TodoDynamoDBService{db: db, table: &todoTable}
}

func (s *TodoDynamoDBService) Create(v api.Entity) error {
	t := v.(*todo.Todo)
	return s.table.Put(t).Run()
}

func (s *TodoDynamoDBService) Get(id [16]byte) (api.Entity, error) {
	todo := &todo.Todo{}
	if err := s.table.Get(partitionKey, id).One(todo); err != nil {
		if err == dynamo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return todo, nil
}

func (s *TodoDynamoDBService) List(_ map[string]string) ([]api.Entity, error) {
	return nil, nil
}

func (s *TodoDynamoDBService) Update(t api.Entity) error {

	return nil
}

func (s *TodoDynamoDBService) Delete(id [16]byte) error {
	return nil
}
