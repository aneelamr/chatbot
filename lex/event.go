package lex

import "github.com/aws/aws-lambda-go/events"

type LexV2Event struct {
	events.LexEvent
	Interpretations []map[string]interface{} `json:"interpretations"`
	SessionState    map[string]interface{}   `json:"sessionState"`
	InputMode       string                   `json:"inputMode"`
}
