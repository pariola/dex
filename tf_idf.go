package main

import (
	"math"
)

// tf calculates the term frequency
func tf(tc, ttc int) float64 {
	return float64(tc) / float64(ttc)
}

// idf calculates the term inverse document frequency
func idf(N, df int32) float64 {
	return math.Log(float64(N) / float64(df))
}

// tf-idf calculates the term frequency, inverse document frequency
func tfIDF(tf, idf float64) float64 {
	return tf * idf
}
