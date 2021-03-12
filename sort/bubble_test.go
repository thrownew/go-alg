package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitBubbleSortSimple(t *testing.T) {
	cases := [][2][]int{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 5, 2, 4, 3}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 2, 3, 1, 1}, []int{1, 1, 2, 2, 3}},
		{[]int{-10, 10, -20, 20, 0}, []int{-20, -10, 0, 10, 20}},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			BubbleSortSimple(c[0])
			assert.Equal(t, c[0], c[1])
		})
	}
}
