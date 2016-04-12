package eventpublisher

import (
	"fmt"
	"testing"
)

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
	fmt.Println("JSON IS:")
	fmt.Println(string(e.GetEventAsJSON()))
}
