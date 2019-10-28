package justify

import (
	"testing"
)

type testcase struct {
	cols     int
	input    string
	expected string
}

func TestCenter(t *testing.T) {
	testcases := []testcase{
		{
			10,
			"Hello",
			"  Hello",
		},
		{
			10,
			"Hello\nWorld",
			"  Hello\n  World",
		},
		{
			10,
			"this is very long text that exceeds columns of the terminal",
			"this is very long text that exceeds columns of the terminal",
		},
		{
			10,
			"日本語",
			"  日本語",
		},
		{
			10,
			"😄🤔",
			"   😄🤔",
		},
		{
			9,
			"odd",
			"   odd",
		},
	}
	for _, c := range testcases {
		cols := c.cols
		input := c.input
		expected := c.expected
		got := Center(cols, input)
		if got != expected {
			t.Errorf("\nexpected:\n%s\ngot:\n%s", expected, got)
		}
	}
}
