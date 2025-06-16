package main

import "fmt"
var gallery = []string{"img1.jpg", "img2.jpg", "img3.jpg", "img4.jpg"}

func challengReslice() {
	fmt.Println("In challengeReslice()")
	fmt.Println("Before:", gallery)
	deleteImage("img2.jpg")
	fmt.Println("After: ", gallery)

	//copy(destination, source)
	//example array := [5]int{1, 2, 3, 4, 5}
	//copy(arr[1:],arr[2:])
	// starting from index 1,starting from index 2
	// copy([2,3,4,5], [3,4,5])
	// shift elements to the left
	// 2 become 3
	// 3 become 4
	// 4 become 5
	// 5 is just 5
	// final copy data will be [1,3,4,5,5]
	array := []int{1, 2, 3, 4, 5}
	copy(array[1:], array[2:])
	fmt.Println("After copy:", array)

}

func deleteImage(filename string) {
	for i := 0; i < len(gallery); i++ {
		if gallery[i] == filename {
			// Shift elements left using copy
			//copy(destination, source)
			//example array := [5]int{1, 2, 3, 4, 5}
			//copy(arr[1:],arr[2:])
			// starting from index 1,starting from index 2
			// copy([2,3,4,5], [3,4,5])
			// shift elements to the left
			// 2 become 3
			// 3 become 4
			// 4 become 5
			// 5 is just 5
			// final copy data will be [1,3,4,5,5]
			copy(gallery[i:], gallery[i+1:])
			// Shrink slice by 1
			gallery = gallery[:len(gallery)-1]
			fmt.Printf("Image '%s' deleted successfully.\n", filename)
			return
		}
	}
	fmt.Printf("Image '%s' not found.\n", filename)
}

