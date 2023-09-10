package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-cmp/cmp"
)

func TestLexEventHandler(t *testing.T) {

	cases := []struct {
		testName  string
		testEvent events.LexEvent
		expected  events.LexResponse
	}{
		{
			testName: "session attributes present",
			testEvent: events.LexEvent{
				CurrentIntent: &events.LexCurrentIntent{
					Name: "TestIntent",
				},
				SessionAttributes: events.SessionAttributes{
					"key1": "value1",
					"key2": "value2",
				},
			},
			expected: events.LexResponse{
				SessionAttributes: events.SessionAttributes{
					"key1": "value1",
					"key2": "value2",
				},
				DialogAction: events.LexDialogAction{
					Type: "Close",
					Message: map[string]string{
						"content":     "Hello from AWS Lambda!",
						"contentType": "PlainText",
					},
					FulfillmentState: "Fulfilled",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			actual, err := LexEventHandler(context.Background(), c.testEvent)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if cmp.Equal(actual, c.expected) {
				t.Errorf("Expected: %v, got: %v", c.expected, actual)
			}
		})
	}
}
