// Package quicksort provides a quicksort implementation
package quicksort

// Sort sorts a slice of int values with the quick sort algorithm.
func Sort(collection []int) {
	quickSort(collection, 0, len(collection)-1)
}

func quickSort(collection []int, low int, high int) {
	if low < high {
		p := partition(collection, low, high)
		quickSort(collection, low, p-1)
		quickSort(collection, p+1, high)
	}
}

func partition(collection []int, low int, high int) int {
	pivot := collection[high]
	i := low
	for j := low; j < high; j++ {
		if collection[j] < pivot {
			temp := collection[j]
			collection[j] = collection[i]
			collection[i] = temp
			i++
		}
	}
	collection[high] = collection[i]
	collection[i] = pivot
	return i
}
