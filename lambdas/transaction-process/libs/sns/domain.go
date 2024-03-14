package sns

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Records struct {
	bucketName string
	objectKey  string
}

func NewRecords(bucketName, objectKey string) *Records {
	return &Records{
		bucketName: bucketName,
		objectKey:  objectKey,
	}
}

func (r *Records) GetBucketName() string {
	return r.bucketName
}
func (r *Records) GetObjectKey() string {
	return r.objectKey
}

func (r *Records) ToString() string {
	return "BucketName: " + r.bucketName + " ObjectKey: " + r.objectKey
}

type CallbackFunc func(ctx context.Context, message events.S3Event, payload []*Records) error
