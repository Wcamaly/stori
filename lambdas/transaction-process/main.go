package main

import (
	"context"
	"os"
	"stori/transaction-lambda/libs/process"
	"stori/transaction-lambda/libs/s3"
	"stori/transaction-lambda/libs/sns"
	"stori/transaction-lambda/libs/transaction"
	"stori/transaction-lambda/libs/user"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func processTransactions(ctx context.Context, message events.S3Event, records []*sns.Records) error {
	s3Service, err := s3.NewS3Service()
	snsClient := sns.NewSnsClientService()
	userService := user.NewUserService(os.Getenv("USER_SERVICE_URL"), 10)

	if err != nil {
		println("Error to create s3 service", err.Error())
		return err
	}

	for _, record := range records {
		if (record.GetBucketName() == "") || (record.GetObjectKey() == "") {
			println("Error to get bucket name or object key empty")
			return nil
		}

		bucketObject, err := s3Service.GetObject(record.GetBucketName(), record.GetObjectKey())
		if err != nil {
			println("Error to get object from s3", err.Error())
			return err
		}

		transactions, err := process.ProcessData(bucketObject, userService)
		if err != nil {
			println("Error to process data", err.Error())
			return err
		}
		transactionService := transaction.NewTransactionService(os.Getenv("TRANSACTION_SERVICE_URL"), 0)
		data, err := transactionService.CreateListTransaction(transactions)

		if err != nil {
			println("Error to create transaction", err.Error())
			return err
		}
		if len(data) == len(transactions) {
			println("Transactions created successfully")
		} else {
			println("Error to create transaction")
		}

		users, err := userService.GetUsers()
		if err != nil {
			println("Error to get users", err.Error())
			return err
		}

		snsClient.Emit(ctx, os.Getenv("SNS_TOPIC"), users)

	}

	return nil
}

func main() {
	lambda.Start(func(ctx context.Context, sqsEvent events.SQSEvent) error {
		println("Start transaction process lambda... - ", time.Now().String())
		return sns.Handler(ctx, sqsEvent, processTransactions)
	})

}
