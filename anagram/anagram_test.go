package main

import "testing"

func TestCheck(t *testing.T) {
	cases := []struct {
		a, b string
		expected bool
	}{
		{
			"not",
			"anagram",
			false,
		},
		{
			"test",
			"tset",
			true,
		},
		{
			"joke",
			"ekoi",
			false,
		},
		{
			"joke",
			"ekoj",
			true,
		},
	}

	for _, c := range cases {
		got := check(c.a, c.b)
		if got != c.expected {
			t.Errorf("check(%+v, %+v) == %+v, should be %+v", c.a, c.b, got, c.expected)
		}
	}
}
