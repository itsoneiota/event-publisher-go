package eventpublisher

import (
	"encoding/json"
	"fmt"
	"time"
)

// func main() {
// 	t := time.Now().UnixNano() / 1000000
// 	b := Event{
// 		Header: Header{
// 			Type:      "Record",
// 			TimeStamp: t,
// 			Origin:    "GOLANG_EVENT_LOG",
// 		},
// 		Body: "{'message Here'}",
// 	}

// 	v, err := json.Marshal(b)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//fmt.Println(b.getType())
// 	//fmt.Println(string(v))

// }
func BuildEvent(o, t string, b interface{}) *Event {
	E := new(Event)
	E.Header = BuildHeader(o, t)
	E.Body = b
	return E
}

func BuildHeader(o, t string) Header {
	H := Header{
		Type:      t,
		TimeStamp: time.Now().UnixNano() / 1000000,
		Origin:    o,
	}
	return H
}

type Header struct {
	Type      string `json:"type"`
	TimeStamp int64  `json:"timeStamp"`
	Origin    string `json:"origin"`
}
type Event struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

func (e *Event) GetType() string {
	t := e.Header.Type
	return t
}

func (e *Event) GetEventAsJson() []byte {
	v, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return v
}
