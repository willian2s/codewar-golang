package numbertostring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willian2s/codewar-golang/pkg/numbertostring"
)

func TestNumberToString(t *testing.T) {
	assert.Equal(t, "1", numbertostring.NumberToString(1))
	assert.Equal(t, "-1", numbertostring.NumberToString(-1))
}
