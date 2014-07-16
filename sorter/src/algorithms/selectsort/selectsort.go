//selectsort.go
package selectsort

func SelectSort(values []int) {
	n := len(values)
	var i, j, min int
	for i = 0; i < n; i++ {
		min = i
		for j = i; j < n; j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		if i != min {
			values[i], values[min] = values[min], values[i]
		}
	}
}
