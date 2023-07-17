package test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willian2s/codewar-golang/pkg/fake_binary"
)

func TestFakeBinary(t *testing.T) {
	assert.Equal(t, "01011110001100111", fake_binary.FakeBinary("45385593107843568"))
	assert.Equal(t, "101000111101101", fake_binary.FakeBinary("509321967506747"))
	assert.Equal(t, "011011110000101010000011011", fake_binary.FakeBinary("366058562030849490134388085"))
	assert.Equal(t, "01111100", fake_binary.FakeBinary("15889923"))
	assert.Equal(t, "100111001111", fake_binary.FakeBinary("800857237867"))
}
