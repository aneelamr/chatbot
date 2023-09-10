package main

import (
	"context"
	"testing"

	"github.com/aneelamr/chatbot/lex"
	"github.com/google/go-cmp/cmp"
)

func TestLexEventHandler(t *testing.T) {

	cases := []struct {
		testName  string
		testEvent lex.LexV2Event
		expected  lex.LexV2Response
	}{
		{
			testName:  "expected response",
			testEvent: lex.LexV2Event{},
			expected: lex.LexV2Response{
				SessionState: lex.SessionState{
					DialogAction: lex.DialogAction{
						SlotElicitationStyle: "Default",
						Type:                 "Close",
					},
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
