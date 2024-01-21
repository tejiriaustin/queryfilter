package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFields(t *testing.T) {

	testCases := []struct {
		name         string
		key          string
		args         func() interface{}
		fieldFinder  func(interface{}, string) []string
		expected     []string
		expectsError bool
	}{
		{
			name: "Basic object",
			key:  "bson",
			args: func() interface{} {
				return struct {
					Name   string `bson:"name"`
					Age    int    `bson:"age"`
					Status string `bson:"status"`
				}{}
			},
			fieldFinder: func(i interface{}, key string) []string {
				return Fields(i, key)
			},
			expected: []string{"name", "status"},
		},
		{
			name: "Basic Nested object",
			key:  "bson",
			args: func() interface{} {
				return struct {
					Name        string `bson:"name"`
					Age         int    `bson:"age"`
					Status      string `bson:"status"`
					Preferences struct {
						Accessibility string `bson:"accessibility"`
						Occupation    string `bson:"occupation"`
					} `bson:"preferences"`
				}{}
			},
			fieldFinder: func(i interface{}, key string) []string {
				return Fields(i, key)
			},
			expected: []string{"name", "status", "preferences.accessibility", "preferences.occupation"},
		},
		{
			name: "Basic Nested object with provided fields",
			key:  "bson",
			args: func() interface{} {
				return struct {
					Name        string `bson:"name"`
					Age         int    `bson:"age"`
					Status      string `bson:"status"`
					Preferences struct {
						Accessibility string `bson:"accessibility"`
						Occupation    string `bson:"occupation"`
					} `bson:"preferences"`
				}{}
			},
			fieldFinder: func(i interface{}, key string) []string {
				return New(i, key).Fields("nets", "towers")
			},
			expected: []string{"nets", "towers", "name", "status", "preferences.accessibility", "preferences.occupation"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fields := tc.fieldFinder(tc.args(), tc.key)
			if tc.expectsError {
				assert.NotEqual(t, tc.expected, fields)
			}
			assert.Equal(t, tc.expected, fields)
		})
	}
}
