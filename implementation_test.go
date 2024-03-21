package Lab_2

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrefixToPostfix(t *testing.T) {
	tests := []struct {
		prefix   string
		expected string
	}{
		{"+ 2 3", "2 3 +"},
		{"* + 2 3 4", "2 3 + 4 *"},
		{"+ * 1 2 3", "1 2 * 3 +"},
		{"- 7 + 6 / 5 * 4 - 3 + 1 2", "7 6 5 4 3 1 2 + - * / + -"},
		{"+ 5 * 3 - 4 / 2 + 2 - 3 + 4 2", "5 3 4 2 2 3 4 2 + - + / - * +"},
		{"* 5 + 4 + 1 - 4 + 8 * 4 + 2 ^ 2 5", "5 4 1 4 8 4 2 2 5 ^ + * + - + + *"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %s", test.prefix), func(t *testing.T) {
			result, err := PrefixToPostfix(test.prefix)
			if test.expected == "" {
				assert.NotNil(t, err, "Expected error for input: %s", test.prefix)
			} else {
				assert.Nil(t, err, "Unexpected error for input: %s", test.prefix)
				assert.Equal(t, test.expected, result, "Incorrect result for input: %s", test.prefix)
			}
		})
	}
}

func TestPrefixToPostfixInvalidInput(t *testing.T) {
	tests := []struct {
		prefix   string
		expected error
	}{
		{" 2 & + 3", fmt.Errorf("invalid input")},
		{" ", fmt.Errorf("invalid input")},
		{"%", fmt.Errorf("invalid input")},
		{"", fmt.Errorf("invalid input")},
		{"+ 2", fmt.Errorf("invalid expression")},
		{"-+ 2 3", fmt.Errorf("invalid input")},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %s", test.prefix), func(t *testing.T) {
			_, err := PrefixToPostfix(test.prefix)
			assert.NotNil(t, err)
			if assert.Error(t, err) {
				assert.Equal(t, test.expected, err)
			}
		})
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("- 5 * 4 2")
	fmt.Println(res)

	// Output:
	// 5 4 2 * -
}
