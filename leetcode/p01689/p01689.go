package p01689

// Time Complexity: O(n)
// Space Complexity: O(1)
func minPartitions(n string) int {
	maximum := 0
	for i := 0; i < len(n); i++ {
		if int(n[i])-48 > maximum {
			maximum = int(n[i]) - 48
		}
	}

	return maximum
}
