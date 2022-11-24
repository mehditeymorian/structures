package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestLinkedList_Add(t *testing.T) {

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
			al := NewLinkedList[int]()

			al.Add(tt.input...)

			assert.Equal(t, al.GetAll(), tt.result)
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {

	tests := []struct {
		name        string
		initial     []int
		removeIndex int
		result      []int
	}{
		{
			name:        "remove first item",
			initial:     []int{1, 2, 3},
			removeIndex: 0,
			result:      []int{2, 3},
		},
		{
			name:        "remove last item",
			initial:     []int{1, 2, 3},
			removeIndex: 2,
			result:      []int{1, 2},
		},
		{
			name:        "remove middle item",
			initial:     []int{1, 2, 3},
			removeIndex: 1,
			result:      []int{1, 3},
		},
		{
			name:        "remove item out of bound",
			initial:     []int{1, 2, 3},
			removeIndex: 4,
			result:      []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ll := NewLinkedList[int](tt.initial...)

			ll.Remove(tt.removeIndex)

			assert.Equal(t, ll.GetAll(), tt.result)
		})
	}
}
