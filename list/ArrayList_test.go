package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestArrayList_Add(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		result []int
	}{
		{
			name:   "add single item",
			input:  []int{1},
			result: []int{1},
		},
		{
			name:   "add multiple items",
			input:  []int{1, 2, 3},
			result: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			al := NewArrayList[int]()

			al.Add(tt.input...)

			assert.Equal(t, al.data, tt.result)
		})
	}
}
