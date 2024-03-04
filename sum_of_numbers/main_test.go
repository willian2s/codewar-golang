package sumofnumbers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sumofnumbers "github.com/willian2s/codewar-golang/sum_of_numbers"
)

func TestSumOfNumbers(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{0, 1, 1},
		{1, 2, 3},
		{5, -1, 14},
		{505, 4, 127759},
		{321, 123, 44178},
		{0, -1, -1},
		{-50, 0, -1275},
		{-1, -5, -15},
		{-5, -5, -5},
		{-505, 4, -127755},
		{-321, 123, -44055},
		{0, 0, 0},
		{-5, -1, -15},
		{5, 1, 15},
		{-17, -17, -17},
		{17, 17, 17},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, sumofnumbers.GetSum(testCase.a, testCase.b))
	}
}
