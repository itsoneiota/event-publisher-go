package EventPublisher

import (
	"fmt"
	"testing"
)

func TestBuildHeader(t *testing.T) {
	h := BuildHeader("Test SUite", "Something else here")
	//fmt.Println(h.Type)
	if h.Type != "Something else here" {
		fmt.Println("Title not the same as set title1")
		t.Fail()
	}
}

func TestBuildEvent(t *testing.T) {
	e := BuildEvent("origin", "Type", "{'body':'Json Stuff Here'}")
	fmt.Println(e.getType())
}
