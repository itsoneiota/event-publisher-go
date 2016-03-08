# event-publisher-go

To create and send a message:
Create new KinesisTransporter
    
       var n = new(KinesisTransporter)
Create an instance of the Kinesis Service. 

    n.KinesisClient = kinesis.New(session.New(), &aws.Config{Region: aws.String("eu-west-1")})

Note: session.New() pulls credentials from 
    Environment settings 
    ~/.aws/credentials 
    IAM Roles if running on EC2 
In that order. See (http://docs.aws.amazon.com/sdk-for-go/latest/v1/developerguide/configuring-sdk.title.html) for more info

Set the Stream name and optionally the partition count on the KinesisTransporter

    n.SetStreamName("event-queue") //Name of Kinesis Stream
    n.SetPartitions(2) //optionally set the number of partitions.
Create a struct that defines your Event Body

    s := struct {
        Arbitary string
        Content string
        Here string
    }{"Arbitary content",
        "To allow you to",
        "Define a JSON format for Event.Body"}

Build the eventUsing an Origin and Type (and Body as defined above)

    e := BuildEvent("origin", "Type", s)

Publish Event to Transporter (in this case Kinesis)

    err := n.Publish(e)

(catch error and handle it if required.)

For Tests a local instance of Kinesalite needs to be running. (Docker image assumed running on http://192.168.99.100:4567 )
Uncomment   //CreateQueue() on line 13 of Event_test.go to create the queue initially. Can be removed for subsequent tests.


