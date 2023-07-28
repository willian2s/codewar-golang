package numbertostring_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	numbertostring "github.com/willian2s/codewar-golang/pkg/number_to_string"
)

func TestNumberToString(t *testing.T) {
	assert.Equal(t, "1", numbertostring.NumberToString(1))
	assert.Equal(t, "-1", numbertostring.NumberToString(-1))
}
