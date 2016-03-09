package eventpublisher

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

const kinesaliteEndpoint = "http://192.168.99.101:4567"

func TestKinesisTransporterPublish(t *testing.T) {
	var n = new(KinesisTransporter)
	n.KinesisClient = kinesis.New(session.New(), &aws.Config{Region: aws.String("eu-west-1"), Endpoint: aws.String(kinesaliteEndpoint)})
	n.SetStreamName("event-queue")
	//n.SetPartitions(2)
	s := struct {
		Body string
		Head string
		Legs string
	}{"Json Body of message",
		"Head here",
		"Legs here"}

	//svc
	e := BuildEvent("origin", "Type", s)
	err := n.Publish(e)

	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestEventPublisherPublish(t *testing.T) {
	var ep = new(EventPublisher)
	ep.SetEnabled(true)

	var n = new(KinesisTransporter)
	n.KinesisClient = kinesis.New(session.New(), &aws.Config{Region: aws.String("eu-west-1"), Endpoint: aws.String(kinesaliteEndpoint)})
	n.SetStreamName("event-queue")
	//n.SetPartitions(2)
	ep.SetTransporter(n)

	s := struct {
		Body string
		Head string
		Legs string
	}{"Json Body of message",
		"Head here",
		"Legs here"}

	//svc
	e := BuildEvent("origin", "Type", s)
	err := ep.Publish(e)

	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}
