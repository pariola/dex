package main

import (
	"math"
	"strings"

	"dex/pkg/tokenizer"
)

// Index
type Index struct {

	// number of documents in the index
	N int

	// dictionary
	d map[string]*PostingList
}

// NewIndex creates a new index instance
func NewIndex() *Index {
	return &Index{
		d: make(map[string]*PostingList),
	}
}

// tfIDF calculates the term inverse document frequency
func (i Index) tfIDF(df int32) float64 {
	return math.Log(float64(i.N) / float64(df))
}

// Add prepare a document for the index
func (i *Index) Add(d Document) {

	i.N++

	t := tokenizer.NewTokenizer(
		strings.NewReader(d.Content),
	)

	for term, indexes := range t.Scan() {

		n := &Node{
			doc:     d.Id,
			indexes: indexes,
		}

		if _, ok := i.d[term]; !ok {
			i.d[term] = new(PostingList)
		}

		p := i.d[term]
		p.Insert(n)

		// update idf
		p.idf = i.tfIDF(p.Size())
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
