package eventpublisher

import (
	"fmt"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type KinesisTransporter struct {
	Message       string
	Transporter   Transporter
	StreamName    string
	KinesisClient *kinesis.Kinesis
	Partitions    int32
}

// Publish publishes event
func (k *KinesisTransporter) Publish(e *Event) {
	k.Message = e.GetType()
	fmt.Println(k.Message)

	params := k.BuildKinesisParams(e)

	resp, err := k.KinesisClient.PutRecord(&params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

}

func (k *KinesisTransporter) GetMessage() string {
	m := k.Message
	fmt.Println("m=" + m)
	return m
}

func (k *KinesisTransporter) GetStreamName() string {
	s := k.StreamName
	return s
}

func (k *KinesisTransporter) SetStreamName(sn string) {
	k.StreamName = sn
}

func (k *KinesisTransporter) SetPartitions(p int32) {
	k.Partitions = p
}

func (k *KinesisTransporter) GetPartition() int32 {
	n := rand.Int31n(k.Partitions)
	return n
}

func (k *KinesisTransporter) BuildKinesisParams(e *Event) kinesis.PutRecordInput {
	b := e.GetEventAsJson()
	fmt.Println(string(b))
	params := kinesis.PutRecordInput{
		Data:         e.GetEventAsJson(),                   // Required
		PartitionKey: aws.String(string(k.GetPartition())), // Required needs to be a string though
		StreamName:   aws.String(k.GetStreamName()),        // Convert the Event struct to JSON
	}

	return params
}
