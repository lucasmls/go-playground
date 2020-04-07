package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	Name          string
	ExpectedCalls []string
	Input         interface{}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	testCases := []testCase{
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

		{
			Name:          "Struct with non string field",
			ExpectedCalls: []string{"Lucas"},
			Input: struct {
				Name string
				Age  int
			}{Name: "Lucas", Age: 20},
		},

		{
			Name:          "Nested fields",
			ExpectedCalls: []string{"Lucas", "Toronto"},
			Input: Person{
				Name: "Lucas",
				Profile: Profile{
					Age:  20,
					City: "Toronto",
				},
			},
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
