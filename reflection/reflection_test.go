package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Mjishu"},
			[]string{"Mjishu"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Mjishu", "London"},
			[]string{"Mjishu", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Mjishu", 19},
			[]string{"Mjishu"},
		},
		{
			"struct with neseted structs",
			Person{"Mjishu", Profile{19, "London"}},
			[]string{"Mjishu", "London"},
		},
		{
			"struct with pointer field",
			&Person{"Mjishu", Profile{19, "London"}},
			[]string{"Mjishu", "London"},
		},
		{
			"struct with slices",
			[]Profile{
				{19, "London"},
				{20, "Tucson"},
			},
			[]string{"London", "Tucson"},
		},
		{
			"struct with arrays",
			[2]Profile{
				{19, "London"},
				{20, "Tucson"},
			}, []string{"London", "Tucson"},
		},
		{
			"struct with map",
			map[string]string{
				"Cow": "Moo",
				"Cat": "Meow",
			},
			[]string{"Moo", "Meow"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
