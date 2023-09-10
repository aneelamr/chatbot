package main

import (
	"context"
	"fmt"

	"github.com/aneelamr/chatbot/lex"
	"github.com/aws/aws-lambda-go/lambda"
)

func LexEventHandler(ctx context.Context, event lex.LexV2Event) (*lex.LexV2Response, error) {
	fmt.Printf("Received an input from Amazon Lex. Current Event: %+v \n", event)

	fmt.Printf("Intent name is: %s \n", event.SessionState.Intent.Name)

	return &lex.LexV2Response{
		SessionState: lex.SessionState{
			DialogAction: lex.DialogAction{
				SlotElicitationStyle: "Default",
				Type:                 "Close",
			},
			Intent: lex.Intent{
				Name:              "OrderFlowers",
				ConfirmationState: "Confirmed",
				State:             "Fulfilled",
			},
		},
		Messages: []lex.Message{
			{
				ContentType: "PlainText",
				Content:     "Thanks for your order! Your Flowers will be delivered in 30 minutes.",
			},
		},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(LexEventHandler)
}
