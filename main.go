package main

import (
	"context"
	"fmt"

	"github.com/aneelamr/chatbot/intents"
	"github.com/aneelamr/chatbot/intents/pizza"
	"github.com/aneelamr/chatbot/lex"
	"github.com/aws/aws-lambda-go/lambda"
)

func LexEventHandler(ctx context.Context, event lex.LexV2Event) (*lex.LexV2Response, error) {
	fmt.Printf("Received an input from Amazon Lex. Current Event: %+v \n", event)

	intent := event.SessionState.Intent.Name
	fmt.Printf("Intent name is: %s \n", intent)

	switch intent {
	case intents.OrderPizza:
		return pizza.OrderHandler(ctx, event)
	default:
		// TODO: Add fallback event handler here.
		return &lex.LexV2Response{}, nil
	}

}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(LexEventHandler)
}
