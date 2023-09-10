package pizza

import (
	"context"

	"github.com/aneelamr/chatbot/lex"
)

func OrderHandler(ctx context.Context, event lex.LexV2Event) (*lex.LexV2Response, error) {
	return &lex.LexV2Response{
		SessionState: lex.SessionState{
			DialogAction: lex.DialogAction{
				SlotElicitationStyle: "Default",
				Type:                 "Close",
				SlotToElicit:         "PizzaType",
			},
			Intent: lex.Intent{
				Name:              "OrderPizza",
				ConfirmationState: "Confirmed",
				State:             "Fulfilled",
			},
		},
		Messages: []lex.Message{
			{
				ContentType: "PlainText",
				Content:     "Thanks for your order! Your Pizza will be delivered in 30 minutes.",
			},
		},
	}, nil
}
