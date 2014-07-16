//bubblesort_test.go
package insertsort

import "testing"

func TestInsertSort(t *testing.T) {
	values := []int{5, 3, 4, 1, 2}
	InsertSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("InsertSort() failed. Got", values, "Expected 1 2 3 4 5")
	}
}

func InsertSortSort2(t *testing.T) {
	values := []int{5, 3, 5, 1, 2}
	InsertSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("InsertSort() failed. Got", values, "Expected 1 2 3 5 5")
	}
}

func InsertSortSort3(t *testing.T) {
	values := []int{5}
	InsertSort(values)
	if values[0] != 5 {
		t.Error("InsertSort() failed. Got", values, "Expected 5")
	}
}
