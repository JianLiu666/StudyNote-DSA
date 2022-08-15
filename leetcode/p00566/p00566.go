package p00566

// Time Complexity: O(n)
// Space Complexity: O(n)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}

	res := make([][]int, r)
	mr, mc := 0, 0
	for rr := 0; rr < len(res); rr++ {
		tmp := make([]int, c)
		for rc := 0; rc < len(tmp); rc++ {
			tmp[rc] = mat[mr][mc]
			if mc < len(mat)-1 {
				mc++
			} else {
				mc = 0
				mr++
			}
		}
		res[rr] = tmp
	}

	return res
}
