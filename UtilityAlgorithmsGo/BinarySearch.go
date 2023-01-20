package BinarySearch

func BinarySearch(sortedArray []int, value int) int {
	length := len(sortedArray)
	lower := 0
	upper := length - 1
	returnIndex := -1

	for lower != upper {
		mid := lower + upper/2
		if sortedArray[mid] == value {
			returnIndex = mid
		} else if value > sortedArray[mid] {
			lower = mid + 1
		} else if value < sortedArray[mid] {
			upper = mid - 1
		}
	}

	return returnIndex
}