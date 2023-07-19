package test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willian2s/codewar-golang/pkg/summation"
)

func TestSummation(t *testing.T) {
	examples := [...][2]int{
		{1, 1},
		{8, 36},
		{22, 253},
		{100, 5050},
		{213, 22791},
	}

	for i := 0; i < len(examples); i++ {
		v := examples[i]
		assert.Equal(t, v[1], summation.Summation(v[0]))
	}
}
