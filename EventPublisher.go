package eventpublisher

type EventPublisher struct {
	Transporter
	Enabled bool
}

func (e *EventPublisher) SetTransporter(t Transporter) {
	e.Transporter = t
}

func (e *EventPublisher) GetTransporter() Transporter {
	if e.Transporter != nil {
		return e.Transporter
	}
	return nil
}

func (e *EventPublisher) SetEnabled(eN bool) {
	e.Enabled = eN
}

func (e *EventPublisher) GetEnabled() bool {
	return e.Enabled
}

func (e *EventPublisher) Publish(eV *Event) {
	e.Transporter.Publish(eV)
}
