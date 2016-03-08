package eventpublisher

//EventPublisher - Describes EventPublisher
type EventPublisher struct {
	Transporter
	Enabled bool
}

//SetTransporter - define transporter for EventPublisher. e.g. accepts a Kinesis client
func (e *EventPublisher) SetTransporter(t Transporter) {
	e.Transporter = t
}

//SetEnabled - Specify whether this is enabled / disabled
func (e *EventPublisher) SetEnabled(eN bool) {
	e.Enabled = eN
}

//GetEnabled - Return the enabled flag.
func (e *EventPublisher) GetEnabled() bool {
	return e.Enabled
}

//Publish Method - Proxy to transporter Publish method. returns error
func (e *EventPublisher) Publish(eV *Event) error {
	err := e.Transporter.Publish(eV)
	return err

}
