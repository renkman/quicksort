// Package quicksort provides a quicksort implementation
package quicksort

//import "sync"

// Sort sorts a slice of int values with the quick sort algorithm.
func Sort(collection []int) {
	quickSort(collection, 0, len(collection)-1)
}

// ConcurrentSort sorts a slice of int values with a conrurrent quick sort algorithm
// (naive approach).
func ConcurrentSort(collection []int) {
	concurrentQuickSort(collection, 0, len(collection)-1)
}

func quickSort(collection []int, low int, high int) {
	if low >= high {
		return
	}
	p := partition(collection, low, high)
	quickSort(collection, low, p-1)
	quickSort(collection, p+1, high)
}

func concurrentQuickSort(collection []int, low int, high int) {
	if low >= high {
		return
	}
	p := partition(collection, low, high)
	go concurrentQuickSort(collection, low, p-1)
	go concurrentQuickSort(collection, p+1, high)
}

func partition(collection []int, low int, high int) int {
	pivot := collection[high]
	i := low
	for j := low; j < high; j++ {
		if collection[j] >= pivot {
			continue
		}
		temp := collection[j]
		collection[j] = collection[i]
		collection[i] = temp
		i++

	}
	collection[high] = collection[i]
	collection[i] = pivot
	return i
}
