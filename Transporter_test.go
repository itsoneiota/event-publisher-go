package eventpublisher

import (
	"fmt"
	"testing"
)

func TestTransporterPublishTest(t *testing.T) {
	n := new(MockTransporter)
	ev := NewMockEvent()
	err := n.Publish(ev)
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
}
