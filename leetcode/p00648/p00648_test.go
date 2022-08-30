package p00648

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	dictionary []string
	sentence   string
}

type output struct {
	ans string
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
				dictionary: []string{"cat", "bat", "rat"},
				sentence:   "the cattle was rattled by the battery",
			},
			output{
				ans: "the cat was rat by the bat",
			},
		},
		{
			input{
				dictionary: []string{"a", "b", "c"},
				sentence:   "aadsfasf absbs bbab cadsfafs",
			},
			output{
				ans: "a a b c",
			},
		},
	}

	for idx, data := range tds {
		ast.Equal(data.o.ans, replaceWords(data.i.dictionary, data.i.sentence), idx+1)
	}
}
