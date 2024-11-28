package internal

import "strings"

type Diff struct{}

func NewDiff() *Diff {
	return &Diff{}
}

func (d Diff) FindLineDiff(a, b string) []string {
	linesA := strings.Split(a, "\n")
	linesB := strings.Split(b, "\n")

	return d.FindLCS(linesA, linesB)
}

// use Hunt-Szymanski algo to find LCS
func (d Diff) FindLCS(linesA, linesB []string) []string {

	positions := make(map[string][]int)

	for i, char := range linesB {
		positions[char] = append(positions[char], i)
	}

	findLIS := func(indices []int) []int {
		lis := []int{}

		for _, index := range indices {
			if len(lis) == 0 || index > lis[len(lis)-1] {
				lis = append(lis, index)
			} else {
				l, r := 0, len(lis)-1
				for l < r {
					mid := (l + r) / 2
					if lis[mid] >= index {
						r = mid
					} else {
						l = mid + 1
					}
				}

				lis[l] = index
			}
		}

		return lis
	}

	indices := []int{}
	for _, char := range linesA {
		if idxList, exists := positions[char]; exists {
			indices = append(indices, idxList...)
		}
	}

	lisIndices := findLIS(indices)

	lcs := make([]string, len(lisIndices))
	for i, idx := range lisIndices {
		lcs[i] = string(linesB[idx])
	}

	return lcs
}
