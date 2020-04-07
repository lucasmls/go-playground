package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedCalls []string
		Input         interface{}
	}{
		{
			Name:          "Struct with one string field.",
			ExpectedCalls: []string{"Lucas"},
			Input: struct {
				Name string
			}{Name: "Lucas"},
		},

		{
			Name:          "Struct with two string fields",
			ExpectedCalls: []string{"Lucas", "Mendes"},
			Input: struct {
				Name    string
				Surname string
			}{Name: "Lucas", Surname: "Mendes"},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
