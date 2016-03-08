package eventpublisher

type Transporter interface {
	Publish(e *Event) error
}
