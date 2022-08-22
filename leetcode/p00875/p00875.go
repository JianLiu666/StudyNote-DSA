package p00875

// Time Complexity: O(logn)
// Space Complexity: O(1)
func minEatingSpeed(piles []int, h int) int {
	low, high := 1, maxInArray(piles)

	for low < high {
		speed := (low + high) / 2
		hours := eating(piles, speed)
		if hours <= h {
			high = speed
		} else {
			low = speed + 1
		}
	}

	return low
}

func eating(piles []int, speed int) (hours int) {
	for _, banara := range piles {
		hours += banara / speed
		if banara%speed != 0 {
			hours++
		}
	}
	return
}

func maxInArray(nums []int) int {
	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
