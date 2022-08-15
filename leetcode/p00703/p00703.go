package p00703

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	NLeft  int
	NRight int
}

func CreateNode(val int) *Node {
	return &Node{
		Val:    val,
		Left:   nil,
		Right:  nil,
		NLeft:  0,
		NRight: 0,
	}
}

type KthLargest struct {
	root *Node
	kth  int
	memo int
	size int
}

func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{
		root: nil,
		kth:  k,
		memo: 0,
		size: 0,
	}

	for _, n := range nums {
		obj.insertBST(n)
	}

	if len(nums) >= k {
		obj.updateKth()
	}

	return obj
}

func (this *KthLargest) Add(val int) int {
	this.insertBST(val)
	if this.memo > val && this.size > this.kth {
		return this.memo
	}

	this.updateKth()
	return this.memo
}

func (this *KthLargest) insertBST(val int) {
	this.size++
	if this.root != nil {
		node := CreateNode(val)
		current := this.root
		for {
			if current.Val >= val {
				current.NLeft++
				if current.Left == nil {
					current.Left = node
					break
				}
				current = current.Left

			} else {
				current.NRight++
				if current.Right == nil {
					current.Right = node
					break
				}
				current = current.Right
			}
		}
	} else {
		this.root = CreateNode(val)
	}
}

func (this *KthLargest) updateKth() {
	current := this.root
	kth := this.kth
	for {
		if current.NRight >= kth {
			current = current.Right
		} else if current.NRight+1 == kth {
			this.memo = current.Val
			break
		} else {
			kth -= current.NRight + 1
			current = current.Left
		}
	}
}
