package s3

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	svc *s3.S3
}

func NewS3Service() (*S3Service, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(os.Getenv("AWS_REGION")),
		Endpoint:         aws.String(os.Getenv("AWS_S3_ENDPOINT")),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL:       aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("AWS_SECRET"), ""),
	})
	if err != nil {
		println("Error create session AWS: ", err.Error())
		return nil, err
	}

	session := s3.New(sess)

	return &S3Service{svc: session}, nil
}

func (s *S3Service) GetObject(bucket, key string) ([]byte, error) {
	result, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(result.Body)
	content := buffer.Bytes()

	return content, nil

}
