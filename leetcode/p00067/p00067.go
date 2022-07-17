package p00067

import (
	"fmt"
	"math/big"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
func addBinary_bruteforce(a string, b string) string {
	ia := len(a) - 1
	ib := len(b) - 1
	carried := 0

	res := ""
	for ia > -1 || ib > -1 || carried > 0 {
		if ia > -1 {
			carried += int(a[ia]) - 48
			ia--
		}
		if ib > -1 {
			carried += int(b[ib]) - 48
			ib--
		}

		res = fmt.Sprintf("%d%s", carried%2, res)
		carried /= 2
	}

	return res
}

func addBinary_builtin(a string, b string) string {
	aInt, bInt, sum := new(big.Int), new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)
	sum.Add(aInt, bInt)

	return sum.Text(2)
}
