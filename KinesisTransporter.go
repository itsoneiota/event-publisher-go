package eventpublisher

import (
	"fmt"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

//KinesisTransporter - Describes KinesisTransporter
type KinesisTransporter struct {
	Message       string
	Transporter   Transporter
	StreamName    string
	KinesisClient *kinesis.Kinesis
	Partitions    int32
}

// Publish publishes event using Kinesis
func (k *KinesisTransporter) Publish(e *Event) error {
	k.Message = e.GetType()
	fmt.Println(k.Message)

	params := k.BuildKinesisParams(e)

	resp, err := k.KinesisClient.PutRecord(&params)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	fmt.Println(resp)
	return nil
}

//GetStreamName - returns StreamName
func (k *KinesisTransporter) GetStreamName() string {
	s := k.StreamName
	return s
}

//SetStreamName - Set Kinesis StreamName
func (k *KinesisTransporter) SetStreamName(sn string) {
	k.StreamName = sn
}

//SetPartitions - Specify the amount of partitions (defaults to 0)
func (k *KinesisTransporter) SetPartitions(p int32) {
	k.Partitions = p
}

//GetPartition - returns either 0 or a random number between 0 and the partion count set above.
func (k *KinesisTransporter) GetPartition() int32 {
	if k.Partitions == 0 {
		return k.Partitions
	}
	n := rand.Int31n(k.Partitions)
	return n
}

//BuildKinesisParams - Build message to send to Kinesis.
func (k *KinesisTransporter) BuildKinesisParams(e *Event) kinesis.PutRecordInput {
	params := kinesis.PutRecordInput{
		Data:         e.GetEventAsJSON(),                   // Required
		PartitionKey: aws.String(string(k.GetPartition())), // Required needs to be a string though
		StreamName:   aws.String(k.GetStreamName()),        // Convert the Event struct to JSON
	}

	return params
}
