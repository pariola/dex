package tokenizer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"
)

const (
	eof = rune(0)
)

type Tokenizer struct {
	p int
	b bytes.Buffer
	r *bufio.Reader
}

func NewTokenizer(r io.Reader) Tokenizer {
	return Tokenizer{
		p: -1,
		r: bufio.NewReader(r),
	}
}

func (t *Tokenizer) read() rune {

	t.p += 1

	r, _, err := t.r.ReadRune()

	if err != nil {
		return eof
	}

	return r
}

func (t *Tokenizer) Scan() map[string][]int {

	terms := make(map[string][]int)

	for {
		r := t.read()

		if r == eof || unicode.IsSpace(r) || unicode.IsPunct(r) {

			if t.b.Len() > 0 {
				term := t.b.String()

				// transformations here
				term = strings.ToLower(term)

				terms[term] = append(terms[term], t.p-t.b.Len())

				// flush buffer
				t.b.Reset()
			}

			if r == eof {
				break
			}

			continue
		}

		t.b.WriteRune(r)
	}

	return terms
}
