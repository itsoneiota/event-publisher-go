package eventpublisher

type eventPublisher interface {
	Publish(eV *Event) error
	SetTransporter(t Transporter)
	SetEnabled(eN bool)
	GetEnabled() bool
}

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

//NewEventPublisher Builds a new event publsher with transporter as param, allowing other transporter types to be passed through.
func NewEventPublisher(kt Transporter, e bool) (svc *EventPublisher, err error) {
	return &EventPublisher{Transporter: kt, Enabled: e}, nil
}

//NewKinesisEventPublisher - Build a publisher with a kinesis client.
//Accepts: s (streamname), e (kinesis endpoint), en (enabled), p (int32 partition count on stream)
func NewKinesisEventPublisher(s, e string, en bool, p int32) (svc *EventPublisher, err error) {
	k := NewKinesisTransporter(s, e, p)
	return &EventPublisher{Transporter: k, Enabled: en}, nil
}
