package iterator

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestArrayList_Iterator(t *testing.T) {
	data := ArrayInt{1, 2, 3, 4, 5}
	iterator := data.Iterator()
	i := 0
	for iterator.HasNext() {
		assert.Equal(t, data[i], iterator.CurrentItem())
		iterator.Next()
		i++
	}
}
