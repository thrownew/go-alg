package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitTree(t *testing.T) {
	tree := NewTree(
		"F",
		NewTree(
			"B",
			NewTree(
				"A",
				nil,
				nil,
			),
			NewTree(
				"D",
				NewTree(
					"C",
					nil,
					nil,
				),
				NewTree(
					"E",
					nil,
					nil,
				),
			),
		),
		NewTree(
			"G",
			nil,
			NewTree(
				"I",
				NewTree(
					"H",
					nil,
					nil,
				),
				nil,
			),
		),
	)

	t.Run("walk pre-order (NLR)", func(t *testing.T) {
		seq := make([]string, 0, 8)
		tree.WalkPreOrder(func(v string, level int) bool {
			seq = append(seq, v)
			return true
		})
		assert.Equal(t, []string{"F", "B", "A", "D", "C", "E", "G", "I", "H"}, seq)
	})

	t.Run("walk in-order (LNR)", func(t *testing.T) {
		seq := make([]string, 0, 8)
		tree.WalkInOrder(func(v string, level int) bool {
			seq = append(seq, v)
			return true
		})
		assert.Equal(t, []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}, seq)
	})

	t.Run("walk post-order (LRN)", func(t *testing.T) {
		seq := make([]string, 0, 8)
		tree.WalkPostOrder(func(v string, level int) bool {
			seq = append(seq, v)
			return true
		})
		assert.Equal(t, []string{"A", "C", "E", "D", "B", "H", "I", "G", "F"}, seq)
	})
}
