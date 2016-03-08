package eventpublisher

type EventPublisherBuilder struct {
	TransporterType string
	Config          EventPublisherConfig
}

type EventPublisherConfig struct {
	TransporterType string
	TransportStream string
	Config          string //JSON BLOB?

}

//Build expects a blob of json here. Can be empty.
func (ep *EventPublisherBuilder) Build(Config EventPublisherConfig) {

}
