package cat_years_dog_years_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	catyearsdogyears "github.com/willian2s/codewar-golang/pkg/cat_years_dog_years"
)

func TestCalculateYears(t *testing.T) {
	assert.Equal(t, [3]int{1, 15, 15}, catyearsdogyears.CalculateYears(1))
	assert.Equal(t, [3]int{2, 24, 24}, catyearsdogyears.CalculateYears(2))
	assert.Equal(t, [3]int{10, 56, 64}, catyearsdogyears.CalculateYears(10))
}
