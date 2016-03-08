# event-publisher-go

To create and send a message:

    var n = new(KinesisTransporter)
    n.KinesisClient = kinesis.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})
    n.SetStreamName("event-queue") //Name of Kinesis Stream
    //n.SetPartitions(2) //optionally set the number of partitions.
    s := struct {
        Body string
        Head string
        Legs string
    }{"Json Body of message",
        "Head here",
        "Legs here"}

    e := BuildEvent("origin", "Type", s)
    err := n.Publish(e)

For Tests a local instance of Kinesalite needs to be running. (Docker image assumed running on http://192.168.99.100:4567 )
Uncomment   //CreateQueue() on line 13 of Event_test.go to create the queue initially. Can be removed for subsequent tests.


