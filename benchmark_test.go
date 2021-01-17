package quicksort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBenchmark_BubbleSort_QuickSort(t *testing.T) {
	var sut [200000]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(sut); i++ {
		sut[i] = rand.Int()
	}
	sutQuick := sut
	sutBubble := sut

	startBubble := time.Now()
	bubbleSort(sutBubble[:])
	stopBubble := time.Now()

	assert.NotEqual(t, sut, sutBubble)
	assert.Equal(t, sut, sutQuick)

	bubbleDuration := stopBubble.Sub(startBubble).Milliseconds()
	fmt.Printf("Bubble sort: %v\n", bubbleDuration)

	startQuick := time.Now()
	Sort(sutQuick[:])
	stopQuick := time.Now()

	quickDuration := stopQuick.Sub(startQuick).Milliseconds()

	fmt.Printf("Quick sort: %v\n", quickDuration)
	assert.Less(t, quickDuration, bubbleDuration)
}
