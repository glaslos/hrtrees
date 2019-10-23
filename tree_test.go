package hrtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree(t *testing.T) {
	r := NewRegularRTree().WithCapacity(3)
	r.Insert(Rect{1, 2, 3, 4}, "A")
	r.Insert(Rect{1, 5, 12, 23}, "B")
	r.Insert(Rect{5, 10, 5, 10}, "C")
	r.Insert(Rect{500, 560, 23, 24}, "D")
	r.Insert(Rect{-200, -100, 50, 10}, "E")
	assert.Equal(t, r.Search(Rect{0, 20, 0, 20}), []interface{}{"A", "C", "B"}, "they should be equal")
}
