package main

func lowerBound(letters []byte, target byte) int {
	left, right := -1, len(letters)
	for left+1 < right {
		mid := left + (right-left)/2
		if letters[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}

func nextGreatestLetter(letters []byte, target byte) byte {
	idx := lowerBound(letters, target+1)
	if idx == len(letters) || letters[idx] < target+1 {
		return letters[0]
	}
	return letters[idx]
}
