//insertsort.go
package insertsort

func InsertSort(values []int) {
	var i, j int
	n := len(values)
	for i = 1; i < n; i++ {
		for j = i; j > 0; j-- {
			if values[j] < values[j-1] {
				values[j], values[j-1] = values[j-1], values[j]
			} else {
				break
			}
		}
	}
}
