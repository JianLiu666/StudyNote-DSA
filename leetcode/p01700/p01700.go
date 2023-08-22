package p01700

// Time Complexity: O(n)
// Space Complexity: O(1)
func countStudents(students []int, sandwiches []int) int {
	count := 0
	for count != len(students) {
		prefer := students[0]
		students = students[1:]

		if prefer == sandwiches[0] {
			sandwiches = sandwiches[1:]
			count = 0
		} else {
			students = append(students, prefer)
			count++
		}
	}

	return len(students)
}
