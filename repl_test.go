package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_clean_input(t *testing.T) {
	var failcount, passcount int
	var green, red, reset = "\033[32m", "\033[31m", "\033[0m"
	var pass, fail = green + "Pass" + reset, red + "Fail" + reset

	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  Hello, World!  ",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "  Hello,    WorLd!  \n",
			expected: []string{"hello,", "world!"},
		},
	}

	for _, test := range tests {
		output := clean_input(test.input)
		if !reflect.DeepEqual(test.expected, output) {
			failcount++
			t.Errorf(`
Input:    %q
Expected: %q
Actual:   %q
%s
`,
				test.input, test.expected, output, fail)
		} else {
			passcount++
			fmt.Printf(`
Input:    %q
Expected: %q
Actual:   %q
%s
`,
				test.input, test.expected, output, pass)
		}
	}
	fmt.Println("----------------------------")
	fmt.Printf("%s count: %s%d%s - ", pass, green, passcount, reset)
	fmt.Printf("%s count: %s%d%s\n", fail, red, failcount, reset)
}
