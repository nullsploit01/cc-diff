package internal

type Diff struct{}

func NewDiff() *Diff {
	return &Diff{}
}

// use Hunt-Szymanski algo to find LCS
func (d Diff) FindLCS(a, b string) string {
	positions := make(map[rune][]int)

	for i, char := range b {
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
	for _, char := range a {
		if idxList, exists := positions[char]; exists {
			indices = append(indices, idxList...)
		}
	}

	lisIndices := findLIS(indices)

	lcs := make([]rune, len(lisIndices))
	for i, idx := range lisIndices {
		lcs[i] = rune(b[idx])
	}

	return string(lcs)
}
