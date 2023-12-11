package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		value    float64
		expected float64
	}{
		{
			value:    -3,
			expected: 3,
		},
		{
			value:    3,
			expected: 3,
		},
		{
			value:    -2.000001,
			expected: 2.000001,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Abs(test.value))
		// if abs := Abs(test.value); abs != test.expected {
		// 	t.Errorf("Ошибка значения, ожидалось %f получено %f", abs, test.expected)
		// }
	}

}

// -3, 3, -2.000001, -0.000000003 и так далее.
