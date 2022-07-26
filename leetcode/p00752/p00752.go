package p00752

// Time Complexity: O(n)
// Space Complexity: O(n)
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	seen := map[string]bool{}
	for _, s := range deadends {
		if s == "0000" {
			return -1
		}
		seen[s] = true
	}

	q := CreateLimitQueue(10000)
	q.Enqueue("0000")
	step := 0
	for !q.IsEmpty() {
		size := q.Size()
		for i := 0; i < size; i++ {
			s, _ := q.Dequeue()

			for i := 0; i < 4; i++ {
				down, up := string(s[i]-1), string(s[i]+1)
				if s[i] == '0' {
					down, up = "9", "1"
				} else if s[i] == '9' {
					down, up = "8", "0"
				}

				wheels := s[:i] + down + s[i+1:]
				if _, exists := seen[wheels]; !exists {
					q.Enqueue(wheels)
					seen[wheels] = true
				}
				wheels = s[:i] + up + s[i+1:]
				if _, exists := seen[wheels]; !exists {
					q.Enqueue(wheels)
					seen[wheels] = true
				}
			}
		}

		step++
		if _, exists := seen[target]; exists {
			return step
		}
	}

	return -1
}

type LimitQueue struct {
	queue    []string
	head     int
	tail     int
	size     int
	capacity int
}

func CreateLimitQueue(capacity int) LimitQueue {
	return LimitQueue{
		queue:    make([]string, capacity),
		head:     0,
		tail:     capacity - 1,
		size:     0,
		capacity: capacity,
	}
}

func (this *LimitQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *LimitQueue) IsFull() bool {
	return this.size == this.capacity
}

func (this *LimitQueue) Size() int {
	return this.size
}

func (this *LimitQueue) Enqueue(wheels string) bool {
	if this.IsFull() {
		return false
	}

	this.tail++
	if this.tail == this.capacity {
		this.tail = 0
	}

	this.queue[this.tail] = wheels
	this.size++

	return true
}

func (this *LimitQueue) Dequeue() (string, bool) {
	if this.IsEmpty() {
		return "", false
	}

	res := this.queue[this.head]

	this.head++
	if this.head == this.capacity {
		this.head = 0
	}

	this.size--
	return res, true
}
