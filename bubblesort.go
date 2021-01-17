package quicksort

func bubbleSort(collection []int) {
	for i := 0; i < len(collection)-1; i++ {
		for j := 0; j < len(collection)-1; j++ {
			if collection[j+1] >= collection[j] {
				continue
			}
			temp := collection[j+1]
			collection[j+1] = collection[j]
			collection[j] = temp
		}
	}
}
