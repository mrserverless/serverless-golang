package services

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/satori/go.uuid"
	"github.com/yunspace/serverless-golang/examples/aws-kinesis/config"
)

// KinesisService is a wrapper around KinesisAPI and config
type KinesisService struct {
	config *config.Config
	client kinesisiface.KinesisAPI
}

// NewKinesisService creates a new KinesisService
func NewKinesisService(cfg *config.Config) *KinesisService {
	return &KinesisService{
		config: cfg,
		client: kinesis.New(session.Must(session.NewSession(
			aws.NewConfig().WithRegion(cfg.Region).WithEndpoint(cfg.KinesisStreamName))),
		),
	}
}

// Publish sends a byte array message to Kinesis
func (s *KinesisService) Publish(bytes []byte) (interface{}, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	params := &kinesis.PutRecordInput{
		Data:         bytes,
		PartitionKey: aws.String(id.String()),
		StreamName:   aws.String(s.config.KinesisStreamName),
	}

	log.Printf("Publishing to Kinesis Stream: %s, Partition: %s, Data: %s",
		aws.StringValue(params.StreamName),
		aws.StringValue(params.PartitionKey),
		string(params.Data))
	result, err := s.client.PutRecord(params)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return result, nil
}
