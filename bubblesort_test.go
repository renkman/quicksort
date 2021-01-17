package quicksort

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort_WithNumbers_SortsNumbers(t *testing.T) {
	expected := []int{5, 7, 22, 24, 500, 1887}
	sut := []int{24, 5, 1887, 7, 22, 500}

	bubbleSort(sut)

	assert.Equal(t, expected, sut)
}

func TestBubbleSort_WithRandomNumbers_SortsNumbers(t *testing.T) {
	var sut [2000]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(sut); i++ {
		sut[i] = rand.Int()
	}

	bubbleSort(sut[:])

	var prev int
	for i := range sut {
		if i == 0 {
			prev = i
			continue
		}

		assert.LessOrEqual(t, prev, i)
		prev = i
	}
}
