package services

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/yunspace/serverless-golang/examples/aws-golang-kinesis/config"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/satori/go.uuid"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"log"
)

type KinesisService struct {
	config *config.Config
	client kinesisiface.KinesisAPI
}

func NewKinesisService(cfg *config.Config) *KinesisService {
	return &KinesisService{
		config: cfg,
		client: kinesis.New(session.Must(session.NewSession(aws.NewConfig()))),
	}
}

func (s *KinesisService) Publish(bytes []byte) (interface{}, error) {
	params := &kinesis.PutRecordInput{
		Data:         bytes,
		PartitionKey: aws.String(uuid.NewV4().String()),
		StreamName:   aws.String(s.config.KinesisStreamName),
	}

	log.Printf("Publishing to Kinesis Stream: %s, Partition: %s, Data: %s",
		params.StreamName, params.PartitionKey, string(params.Data))
	result, err := s.client.PutRecord(params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}
