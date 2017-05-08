package main

import "testing"

func TestCheck(t *testing.T) {
	cases := []struct {
		i        int
		expected string
	}{
		{
			1,
			"1",
		},
		{
			3,
			"fizz",
		},
		{
			4,
			"4",
		},
		{
			5,
			"buzz",
		},
		{
			15,
			"fizzbuzz",
		},
	}

	for _, c := range cases {
		got := fizzbuzz(c.i)
		if got != c.expected {
			t.Errorf("fizzbuzz(%+v) == %+v, should be %+v", c.i, got, c.expected)
		}
	}
}
