package main

import "testing"

func TestFizzbuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{9, "Fizz"},
		{10, "Buzz"},
	}
	for _, c := range cases {
		if got := fizzbuzz(c.in); got != c.want {
			t.Errorf("fizzbuzz(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}
