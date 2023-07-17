package cat_years_dog_years_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willian2s/codewar-golang/pkg/cat_years_dog_years"
)

func TestCalculateYears(t *testing.T) {
	assert.Equal(t, [3]int{1, 15, 15}, cat_years_dog_years.CalculateYears(1))
	assert.Equal(t, [3]int{2, 24, 24}, cat_years_dog_years.CalculateYears(2))
	assert.Equal(t, [3]int{10, 56, 64}, cat_years_dog_years.CalculateYears(10))
}
