package main

import (
	"strings"

	"dex/pkg/tokenizer"
)

// Index
type Index struct {

	// number of documents in the index
	N int32

	// dictionary
	d map[string]*PostingList
}

// NewIndex creates a new index instance
func NewIndex() *Index {
	return &Index{
		d: make(map[string]*PostingList),
	}
}

// Add prepare a document for the index
func (i *Index) Add(d Document) {

	i.N++

	t := tokenizer.NewTokenizer(
		strings.NewReader(d.Content),
	)

	// extract terms
	terms := t.Scan()

	// total terms count
	ttc := len(terms)

	for term, indexes := range terms {

		n := &Node{
			doc:     d.Id,
			indexes: indexes,
		}

		n.tf = tf(len(indexes), ttc)

		if _, ok := i.d[term]; !ok {
			i.d[term] = new(PostingList)
		}

		p := i.d[term]
		p.Insert(n)

		// update idf
		p.idf = idf(i.N, p.Size())
	}
}

// TODO
// Search
func (i Index) Search(query string) {

	// t := tokenizer.NewTokenizer(
	// 	strings.NewReader(query),
	// )

	// for term, _ := range t.Scan() {
	// 	p := i.d[term]
	// 	p.
	// }
}
