package eventpublisher

import (
	"fmt"
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

//FirehoseTransporter - Describes FirehoseTransporter
type FirehoseTransporter struct {
	Message       string
	StreamName    string
	KinesisClient *firehose.Firehose
	Partitions    int32
}

// Publish publishes event using Kinesis
func (k *FirehoseTransporter) Publish(e *Event) error {
	k.Message = e.GetType()

	params := k.BuildFirehoseParams(e)

	resp, err := k.KinesisClient.PutRecord(&params)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	if resp == nil {

	}

	return nil
}

//GetStreamName - returns StreamName
func (k *FirehoseTransporter) GetStreamName() string {
	s := k.StreamName
	return s
}

//SetStreamName - Set Kinesis StreamName
func (k *FirehoseTransporter) SetStreamName(sn string) {
	k.StreamName = sn
}

//SetPartitions - Specify the amount of partitions (defaults to 0)
func (k *FirehoseTransporter) SetPartitions(p int32) {
	k.Partitions = p
}

//GetPartition - returns either 0 or a random number between 0 and the partion count set above.
func (k *FirehoseTransporter) GetPartition() int32 {
	if k.Partitions == 0 {
		return k.Partitions
	}
	n := rand.Int31n(k.Partitions)
	return n
}

//BuildFirehoseClient Builds a kinesis client, allowing specification of endpoint (used for when running locally)
func BuildFirehoseClient(endpoint string) *firehose.Firehose {
	if endpoint != "" {
		return firehose.New(session.New(), &aws.Config{Region: aws.String("eu-west-1"), Endpoint: aws.String(endpoint)})
	}
	return firehose.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

}

//BuildFirehoseParams - Build message to send to Kinesis.
func (k *FirehoseTransporter) BuildFirehoseParams(e *Event) firehose.PutRecordInput {
	params := firehose.PutRecordInput{
		DeliveryStreamName: aws.String(k.GetStreamName()), // Convert the Event struct to JSON
		Record: &firehose.Record{
			Data: e.GetEventAsJSON(),
		}, // Required

	}

	return params
}

/*NewFirehoseTransporter  Build a kinesis transporter
/ s: Kinesis Stream Name
/ e: Kinesis Endpoint (use empty string if not locally run)
/ p: Partition count on kinesis stream
/ Returns *KinesisTransporter
*/
func NewFirehoseTransporter(s, e string, p int32) (k *FirehoseTransporter) {
	return &FirehoseTransporter{
		Message:       "",
		StreamName:    s,
		KinesisClient: BuildFirehoseClient(e),
		Partitions:    p,
	}
}
