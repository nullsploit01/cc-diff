package internal

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestFindLineDiff(t *testing.T) {
	cmd := &cobra.Command{}
	out := &bytes.Buffer{}
	cmd.SetOut(out)

	d := NewDiff(cmd)

	a := `Line 1
Line 2
Line 3`
	b := `Line 1
Line 3
Line 4`

	d.FindLineDiff(a, b)

	expected := `< Line 2
> Line 4
`

	if out.String() != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, out.String())
	}
}

func TestFindLCS(t *testing.T) {
	cmd := &cobra.Command{}
	d := NewDiff(cmd)

	linesA := []string{"Line 1", "Line 2", "Line 3"}
	linesB := []string{"Line 1", "Line 3", "Line 4"}

	expectedLCS := []string{"Line 1", "Line 3"}

	lcs := d.FindLCS(linesA, linesB)

	if len(lcs) != len(expectedLCS) {
		t.Errorf("expected LCS length %d, got %d", len(expectedLCS), len(lcs))
	}

	for i, line := range expectedLCS {
		if lcs[i] != line {
			t.Errorf("expected LCS[%d] = %s, got %s", i, line, lcs[i])
		}
	}
}

func TestFindFileDiff(t *testing.T) {
	cmd := &cobra.Command{}
	out := &bytes.Buffer{}
	cmd.SetOut(out)

	d := NewDiff(cmd)

	file1, err := os.CreateTemp("", "file1")
	if err != nil {
		t.Fatalf("failed to create temp file1: %v", err)
	}
	defer os.Remove(file1.Name())

	file2, err := os.CreateTemp("", "file2")
	if err != nil {
		t.Fatalf("failed to create temp file2: %v", err)
	}
	defer os.Remove(file2.Name())

	file1Content := `Line 1
Line 2
Line 3`
	file2Content := `Line 1
Line 3
Line 4`

	file1.WriteString(file1Content)
	file1.Close()

	file2.WriteString(file2Content)
	file2.Close()

	err = d.FindFileDiff(file1.Name(), file2.Name())
	if err != nil {
		t.Fatalf("FindFileDiff failed: %v", err)
	}

	expected := `< Line 2
> Line 4
`

	if out.String() != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, out.String())
	}
}
