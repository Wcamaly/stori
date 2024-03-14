package main

import (
	"context"
	"os"
	"stori/email-service/email"
	"stori/email-service/sns"
	"stori/email-service/transaction"
	"stori/email-service/user"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func ProvessEmails(ctx context.Context, message events.SNSEntity, payload interface{}) error {
	listUsers := user.ProcessJSONData(payload)
	tran := transaction.NewTransactionService(os.Getenv("TRANSACTION_SERVICE_URL"), 0)
	emailService := email.NewEmailService()

	for _, user := range listUsers {
		movements, err := tran.GetMovementByUserId(user.ID)
		if err != nil {
			println("Error to get movements: ", err.Error())
			return err
		}

		balance, err := tran.GetBalanceByUserId(user.ID)
		if err != nil {
			println("Error to get balanc: ", err.Error())
			return err
		}

		println("Balance: ", balance.Balance, " - ", balance.Credit, " - ", balance.Debit, " - ", balance.UserId)

		for _, movement := range movements {
			println("Movement: ", movement.UserID, " - ", movement.Month, " - ", movement.Increment, " - ", movement.Decrement)
		}

		emailContent, err := email.ParseEmail(movements, balance, user)
		if err != nil {
			println("Error parsing email content: ", err.Error())
			return err
		}

		//TODO Subject should be a constant
		emailService.SendEmail(*user, "Your Balance", emailContent)

	}

	return nil
}

func main() {
	lambda.Start(func(ctx context.Context, snsEvent events.SNSEvent) error {
		println("Starting email-process lambda... - ", time.Now().String())
		return sns.HandleRequest(ctx, snsEvent, ProvessEmails)
	})

}
