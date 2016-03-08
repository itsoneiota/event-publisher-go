package eventpublisher

import (
	"encoding/json"
	"fmt"
	"time"
)

/*BuildEvent - Accepts origin and type as strings, and an interface for body
Example call:
body := struct {
	Body string
	RequestURI string
	ResponseCode string
}{"Json Body of message",
	"requestURI Here",
	"ResponseCode Here"}

e := BuildEvent("origin", "Type", body)

Build Header from Origin and Type, and returns an Event
*/
func BuildEvent(o, t string, b interface{}) *Event {
	E := new(Event)
	E.Header = BuildHeader(o, t)
	E.Body = b
	return E
}

/*BuildHeader - called form BuildEvent function.
Generates a valide Event.Header from Origin and Type
Inserts Millisecond timestamp
*/
func BuildHeader(o, t string) Header {
	H := Header{
		Type:      t,
		TimeStamp: time.Now().UnixNano() / 1000000,
		Origin:    o,
	}
	return H
}

/*Header - Struct
Contains Header information common to all Events
*/
type Header struct {
	Type      string `json:"type"`
	TimeStamp int64  `json:"timeStamp"`
	Origin    string `json:"origin"`
}

/*Event - Struct
Contains a Header and a Body, body is passed inas an empty interfcae here to allow for generating structs on the fly.
*/
type Event struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

/*GetType Returns Event type form Header - Not used in general usage.
 */
func (e *Event) GetType() string {
	t := e.Header.Type
	return t
}

//GetEventAsJSON returns byte array of the Event as JSON Used when sending to Kinesis.
func (e *Event) GetEventAsJSON() []byte {
	v, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return v
}
