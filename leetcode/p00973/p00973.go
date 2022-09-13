package p00973

type Pos struct {
	X        int
	Y        int
	Distance int
}

// Time Complexity: O(nlogn)
// Space Complexity: O(n)
func kClosest(points [][]int, k int) [][]int {
	list := make([]*Pos, len(points))

	end := len(list) - 1
	for i := end; i >= 0; i-- {
		list[i] = &Pos{X: points[i][0], Y: points[i][1], Distance: distance(points[i])}
		shiftDown(list, i, end)
	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		res = append(res, []int{list[0].X, list[0].Y})
		list[0], list[end] = list[end], list[0]
		end--
		shiftDown(list, 0, end)
	}

	return res
}

func distance(pos []int) int {
	return pos[0]*pos[0] + pos[1]*pos[1]
}

func shiftDown(list []*Pos, start, end int) {
	parent := start
	for {
		child := parent*2 + 1
		if child > end {
			break
		}

		if child+1 <= end && list[child+1].Distance < list[child].Distance {
			child++
		}

		if list[child].Distance < list[parent].Distance {
			list[child], list[parent] = list[parent], list[child]
			parent = child
		} else {
			break
		}
	}
}
