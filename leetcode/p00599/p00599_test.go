package p00599

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	list1 []string
	list2 []string
}

type output struct {
	ans []string
}

type testdata struct {
	i input
	o output
}

func TestQuestion(t *testing.T) {
	ast := assert.New(t)

	tds := []testdata{
		{
			input{
				list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			},
			output{
				ans: []string{"Shogun"},
			},
		},
		{
			input{
				list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"KFC", "Shogun", "Burger King"},
			},
			output{
				ans: []string{"Shogun"},
			},
		},
		{
			input{
				list1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				list2: []string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
			},
			output{
				ans: []string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
			},
		},
	}

	for _, data := range tds {
		ast.Equal(data.o.ans, findRestaurant(data.i.list1, data.i.list2))
	}
}
