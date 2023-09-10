package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func LexEventHandler(ctx context.Context, event events.LexEvent) (*events.LexResponse, error) {
	fmt.Printf("Received an input from Amazon Lex. Current Event: %+v", event)

	messageContent := "Hello from AWS Lambda!"

	return &events.LexResponse{
		SessionAttributes: event.SessionAttributes,
		DialogAction: events.LexDialogAction{
			Type: "Close",
			Message: map[string]string{
				"content":     messageContent,
				"contentType": "PlainText",
			},
			FulfillmentState: "Fulfilled",
		},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(LexEventHandler)
}
