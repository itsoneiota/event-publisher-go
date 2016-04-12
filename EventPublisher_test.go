package eventpublisher

import (
	"fmt"
	"testing"
)

const kinesaliteEndpoint = "http://192.168.99.101:4567"

func TestEventPublisherSetEnabled(t *testing.T) {
	e := new(EventPublisher)
	e.SetEnabled(true)
	if e.Enabled != true {
		fmt.Println("Publisher Not enabled")
		t.Fail()
	}
}

func TestEventPublisherSetDisabled(t *testing.T) {
	e := new(EventPublisher)
	e.SetEnabled(false)
	if e.Enabled != false {
		fmt.Println("Publisher Not disabled")
		t.Fail()
	}
}

func TestEventPublisherSetTransporter(t *testing.T) {
	e := new(EventPublisher)
	n := new(MockTransporter)
	e.SetTransporter(n)

}
func TestEventPublisherPublish(t *testing.T) {
	e := new(EventPublisher)
	n := new(MockTransporter)
	ev := NewMockEvent()
	e.SetTransporter(n)
	err := e.Publish(ev)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}

func TestNewEventPublisherMethod(t *testing.T) {
	n := new(MockTransporter)
	ev := NewMockEvent()
	en := true
	e, err := NewEventPublisher(n, en)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	fmt.Println(e.GetEnabled())
	er := e.Publish(ev)
	if er != nil {
		fmt.Println(er.Error())
		t.Fail()
	}
}

//Commented out below only really testable on local box.

// func TestNewKinesisEventPublisher(t *testing.T) {
// 	//eval $(docker-machine env development)CreateQueue()
// 	//n := new(MockTransporter)
// 	ev := NewMockEvent()
// 	en := true
// 	e, err := NewKinesisEventPublisher("event-queue", "http://192.168.99.101:4567", en, 1)
// 	if err != nil {
// 		t.Fail()
// 	}
// 	fmt.Println(e.GetEnabled())
// 	er := e.Publish(ev)
// 	if er != nil {
// 		t.Fail()
// 	}
// }

// func CreateQueue() {

// 	KinesisClient := kinesis.New(session.New(), &aws.Config{Region: aws.String("eu-west-1"), Endpoint: aws.String(kinesaliteEndpoint)})

// 	params := &kinesis.CreateStreamInput{
// 		ShardCount: aws.Int64(2),              // Required
// 		StreamName: aws.String("event-queue"), // Required
// 	}
// 	resp, err := KinesisClient.CreateStream(params)

// 	if err != nil {
// 		// Print the error, cast err to awserr.Error to get the Code and
// 		// Message from an error.
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	// Pretty-print the response data.
// 	fmt.Println(resp)

// }
