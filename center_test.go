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
			cols:     10,
			input:    "Hello",
			expected: "  Hello",
		},
		{
			cols:     10,
			input:    "Hello\nWorld",
			expected: "  Hello\n  World",
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
