package sns

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent, callback CallbackFunc) error {
	for _, records := range snsEvent.Records {

		snsRecord := records.SNS
		message := snsRecord.Message

		var payload interface{}
		err := json.Unmarshal([]byte(message), &payload)
		if err != nil {
			fmt.Printf("error to decode payload: %v\n", err.Error())
			continue
		}
		if err := callback(ctx, snsRecord, payload); err != nil {
			println("error to call Callback: %v\n", err.Error())
		}
	}
	return nil
}
