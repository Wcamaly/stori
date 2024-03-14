package sns

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type CallbackFunc func(ctx context.Context, message events.SNSEntity, payload interface{}) error
