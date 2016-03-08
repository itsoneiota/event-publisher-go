package eventpublisher

//Transporter - Defines Transporter Inteerface - All transporters must have the following:
type Transporter interface {
	Publish(e *Event) error
}
