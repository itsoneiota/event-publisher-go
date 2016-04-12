package eventpublisher

import "fmt"

//Transporter - Defines Transporter Inteerface - All transporters must have the following:
type Transporter interface {
	Publish(e *Event) error
}

//MockTransporter type, defines the mock transporter
type MockTransporter struct {
	transporterType string
}

//Publish Mock traporter Publish method. To pass it off as a transporter.
func (t *MockTransporter) Publish(e *Event) error {
	fmt.Println(string(e.GetEventAsJSON()))
	return nil
}
