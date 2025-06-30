package main

func bubbleSort(sl []int) {
	// passes through the slice:
	for pass:=1; pass < len(sl); pass++ {
		// one pass:
		for i:=0; i < len(sl) - pass; i++ {		// the bigger value 'bubbles up' to the last position 
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}
}