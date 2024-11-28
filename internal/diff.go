package internal

import (
	"strings"

	"github.com/spf13/cobra"
)

type Diff struct {
	cmd *cobra.Command
}

func NewDiff(cmd *cobra.Command) *Diff {
	return &Diff{
		cmd: cmd,
	}
}

func (d Diff) FindLineDiff(a, b string) {
	linesA := strings.Split(a, "\n")
	linesB := strings.Split(b, "\n")

	d.PrintDiff(linesA, linesB, d.FindLCS(linesA, linesB))

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

func (d Diff) PrintDiff(linesA, linesB, lcs []string) {
	lcsIndex := 0
	i, j := 0, 0

	for i < len(linesA) || j < len(linesB) {
		isInLCS := lcsIndex < len(lcs) &&
			i < len(linesA) && linesA[i] == lcs[lcsIndex] &&
			j < len(linesB) && linesB[j] == lcs[lcsIndex]

		if isInLCS {
			i++
			j++
			lcsIndex++
		} else {
			if i < len(linesA) && (lcsIndex >= len(lcs) || linesA[i] != lcs[lcsIndex]) {
				d.cmd.OutOrStdout().Write([]byte(("< " + linesA[i]) + "\n"))
				i++
			}
			if j < len(linesB) && (lcsIndex >= len(lcs) || linesB[j] != lcs[lcsIndex]) {
				d.cmd.OutOrStdout().Write([]byte(("> " + linesB[j]) + "\n"))
				j++
			}
		}
	}
}
