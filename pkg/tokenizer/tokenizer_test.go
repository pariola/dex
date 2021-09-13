package tokenizer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer_Scan(t *testing.T) {

	z := NewTokenizer(
		strings.NewReader("i am a boy"),
	)

	assert.NotNil(t, z)

	terms := z.Scan()

	assert.Len(t, terms, 4)

	assert.Equal(t, terms["i"], []int{0})
	assert.Equal(t, terms["am"], []int{2})
	assert.Equal(t, terms["a"], []int{5})
	assert.Equal(t, terms["boy"], []int{7})
}
