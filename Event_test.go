package eventpublisher

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func TestBuildHeader(t *testing.T) {
	//CreateQueue()
	h := BuildHeader("Origin", "EventType")
	//fmt.Println(h.Type)
	if h.Type != "EventType" {
		fmt.Println("Event type is not = EventType")
		t.Fail()
	} else {
		fmt.Println("Event Type == EventType")
	}
}

func TestBuildEvent(t *testing.T) {
	s := struct {
		Body string
		Head string
		Legs string
	}{"Json Body of message",
		"Head here",
		"Legs here"}

	e := BuildEvent("origin", "Type", s)
	if e.GetType() != "Type" {
		fmt.Println("Type is not what we sent through")
		t.Fail()
	} else {
		fmt.Println("Event type = Type")
	}
}
func CreateQueue() {

	KinesisClient := kinesis.New(session.New(), &aws.Config{Region: aws.String("eu-west-1"), Endpoint: aws.String("http://192.168.99.100:4567")})

	params := &kinesis.CreateStreamInput{
		ShardCount: aws.Int64(2),              // Required
		StreamName: aws.String("event-queue"), // Required
	}
	resp, err := KinesisClient.CreateStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)

}
