package main

import (
	"context"
	"fmt"

	"github.com/aneelamr/chatbot/lex"
	"github.com/aws/aws-lambda-go/lambda"
)

func LexEventHandler(ctx context.Context, event lex.LexV2Event) (*lex.LexV2Response, error) {
	fmt.Printf("Received an input from Amazon Lex. Current Event: %+v", event)

	//messageContent := "Hello from AWS Lambda!"
	return &lex.LexV2Response{
		SessionState: lex.SessionState{
			DialogAction: lex.DialogAction{
				SlotElicitationStyle: "Default",
				Type:                 "Close",
			},
		},
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(LexEventHandler)
}
