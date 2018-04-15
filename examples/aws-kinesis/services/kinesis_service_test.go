package services

import (
	"io/ioutil"
	"testing"

	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yunspace/serverless-golang/examples/aws-kinesis/config"
)

func TestDataPublisher(t *testing.T) {
	// given
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockKinesis := NewMockKinesisAPI(mockCtrl)
	svc := &KinesisService{
		config.NewConfig(),
		mockKinesis,
	}
	bytes, _ := ioutil.ReadFile("event.json")

	// when
	mockKinesis.EXPECT().PutRecord(gomock.Any()).Return(&kinesis.PutRecordOutput{}, nil)
	result, err := svc.Publish(bytes)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
