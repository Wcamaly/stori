package sns

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func Handler(ctx context.Context, sqsEvent events.SQSEvent, callback CallbackFunc) error {

	for _, message := range sqsEvent.Records {

		var s3Event events.S3Event
		err := json.Unmarshal([]byte(message.Body), &s3Event)
		if err != nil {
			println("Error deserializando el evento S3 del body del mensaje SQS: ", err)
			return err
		}
		payload := make([]*Records, 0)
		for _, record := range s3Event.Records {
			if record.S3.Bucket.Name != "" || record.S3.Object.Key != "" {
				record := NewRecords(record.S3.Bucket.Name, record.S3.Object.Key)
				payload = append(payload, record)
			}
		}
		if err := callback(ctx, s3Event, payload); err != nil {
			println("error al procesar el mensaje: ", err)
		}
	}
	return nil
}

type SnsClientService struct {
	snsClient *sns.SNS
}

func NewSnsClientService() *SnsClientService {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(os.Getenv("AWS_REGION")),
		Endpoint:         aws.String(os.Getenv("AWS_ENDPOINT")),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),

		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("AWS_SECRET"), ""),
	})
	if err != nil {
		println("Error creating AWS session: ", err.Error())
		return nil
	}

	snsc := sns.New(sess)

	return &SnsClientService{
		snsClient: snsc,
	}
}

func (s *SnsClientService) Emit(ctx context.Context, topicArn string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		println("Error to marshal payload: ", err.Error())
		return err
	}
	println("Payload marshaled", string(data))
	_, err = s.snsClient.PublishWithContext(ctx, &sns.PublishInput{
		TopicArn: &topicArn,
		Message:  aws.String(string(data)),
	})
	if err != nil {
		println("Error to publish message: ", err.Error())
		return err
	}
	return nil
}
