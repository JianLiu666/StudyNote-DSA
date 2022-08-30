package p00421

// Time Complexity: O(mn), where m is the number of nums, n is the height of trie
// Space Complexity: O(2^n), where n is the height of trie
func findMaximumXOR(nums []int) int {
	height := lengthOfBinary(max(nums))

	trie := CreateTrie(height)
	for _, n := range nums {
		trie.Put(n)
	}

	maximum := 0
	for _, n := range nums {
		xor := n ^ trie.FindMaximumXOR(n)
		if xor > maximum {
			maximum = xor
		}
	}
	return maximum
}

// 找出陣列中最大的元素
func max(nums []int) int {
	maximum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maximum {
			maximum = nums[i]
		}
	}
	return maximum
}

// 將數字轉成二進制後的數字長度
func lengthOfBinary(num int) int {
	count := 0
	for num > 0 {
		num >>= 1
		count++
	}
	return count
}

type Node struct {
	Childs [2]*Node
	Val    int
}

func CreateNode() *Node {
	return &Node{
		Childs: [2]*Node{},
		Val:    0,
	}
}

type Trie struct {
	Root   *Node
	Height int // Trie 可能的最大高度
}

func CreateTrie(height int) Trie {
	return Trie{
		Root:   CreateNode(),
		Height: height,
	}
}

// Time Complexity: O(n), where n ihe height of trie
// Space Complexity: O(n)
func (t *Trie) Put(val int) {
	current := t.Root
	for i := t.Height - 1; i >= 0; i-- {
		b := (val >> i) & 1
		if current.Childs[b] == nil {
			current.Childs[b] = CreateNode()
		}
		current = current.Childs[b]
	}
	current.Val = val
}

// Time Complexity: O(n), where n ihe height of trie
// Space Complexity: O(1)
func (t *Trie) FindMaximumXOR(val int) int {
	current := t.Root
	for i := t.Height - 1; i >= 0; i-- {
		b := (val >> i) & 1
		if b == 0 && current.Childs[1] != nil {
			current = current.Childs[1]
		} else if b == 1 && current.Childs[0] != nil {
			current = current.Childs[0]
		} else {
			current = current.Childs[b]
		}
	}
	return current.Val
}
