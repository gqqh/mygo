//bubblesort.go
package bubblesort

func BubbleSort(values []int) {
	flag := true
	n := len(values)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		} //end for j
		if flag == true {
			break
		}
	} //end for i
}
