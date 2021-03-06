package player

import "math"

func computeScore(keywords1, keywords2 []string) float64 {
	var total float64 = 0
	var seen map[string]bool = make(map[string]bool)

	for _, k1 := range keywords1 {
		if !seen[k1] {
			for _, k2 := range keywords2 {
				d := inverseLevenshteinDistance(k1, k2)
				if math.IsInf(d, 0) {
					total = total + 10
				} else {
					total = total + d
				}
			}
			seen[k1] = true
		}
	}

	return total
}

func inverseLevenshteinDistance(word1, word2 string) float64 {
	// Code from http://rosettacode.org/wiki/Levenshtein_distance#Go
	d := make([][]int, len(word1)+1)
	for i := range d {
		d[i] = make([]int, len(word2)+1)
	}
	for i := range d {
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}
	for j := 1; j <= len(word2); j++ {
		for i := 1; i <= len(word1); i++ {
			if word1[i-1] == word2[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				min := d[i-1][j]
				if d[i][j-1] < min {
					min = d[i][j-1]
				}
				if d[i-1][j-1] < min {
					min = d[i-1][j-1]
				}
				d[i][j] = min + 1
			}
		}

	}
	return 1 / float64(d[len(word1)][len(word2)])
}
