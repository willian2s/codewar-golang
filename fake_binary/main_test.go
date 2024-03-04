package fakebinary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	fakebinary "github.com/willian2s/codewar-golang/fake_binary"
)

func TestFakeBinary(t *testing.T) {
	assert.Equal(t, "01011110001100111", fakebinary.FakeBinary("45385593107843568"))
	assert.Equal(t, "101000111101101", fakebinary.FakeBinary("509321967506747"))
	assert.Equal(t, "011011110000101010000011011", fakebinary.FakeBinary("366058562030849490134388085"))
	assert.Equal(t, "01111100", fakebinary.FakeBinary("15889923"))
	assert.Equal(t, "100111001111", fakebinary.FakeBinary("800857237867"))
}
