//qsort.go
package qsort

func quickSort(values []int, left, right int) {
	temp := values[left]
	i, j := left, right

	if len(values) < 2 {
		return
	}
	for i < j {
		for j > i && values[j] >= temp { // <--j
			j--
		}
		values[i] = values[j] //values[i] = values[j]< temp

		for i < j && values[i] <= temp { // i -->
			i++
		}
		values[j] = values[i] //values[j] = values[i] > temp
	} //end for
	//i == j
	values[i] = temp
	if i-left > 1 {
		quickSort(values, left, i-1)
	}
	if right-j > 1 {
		quickSort(values, j+1, right)
	}
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}
