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
	t.Run("generic types", func(t *testing.T) {
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

			{
				Name:          "Pointer to things",
				ExpectedCalls: []string{"Lucas", "Toronto"},
				Input: &Person{
					Name: "Lucas",
					Profile: Profile{
						Age:  20,
						City: "Toronto",
					},
				},
			},

			{
				Name:          "Slices",
				ExpectedCalls: []string{"Toronto", "North Vancouver"},
				Input: []Profile{
					{Age: 20, City: "Toronto"},
					{Age: 21, City: "North Vancouver"},
				},
			},

			{
				Name:          "Arrays",
				ExpectedCalls: []string{"Toronto", "North Vancouver"},
				Input: [2]Profile{
					{Age: 20, City: "Toronto"},
					{Age: 21, City: "North Vancouver"},
				},
			},

			{
				Name:          "Maps",
				ExpectedCalls: []string{"Bar", "Foo"},
				Input: map[string]string{
					"Foo": "Bar",
					"Bar": "Foo",
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
	})

	t.Run("works nicely with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, receveid []string, expected string) {
	contains := false
	for _, x := range receveid {
		if x == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", receveid, expected)
	}
}
