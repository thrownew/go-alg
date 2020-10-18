package list

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitList(t *testing.T) {
	t.Run("new predefined", func(t *testing.T) {
		l := NewList("1", "2", "3")
		assert.Equal(t, 3, l.Len())
		assert.False(t, l.Contain(""))
		assert.True(t, l.Contain("1"))
		assert.True(t, l.Contain("2"))
		assert.True(t, l.Contain("3"))
	})
	t.Run("new empty", func(t *testing.T) {
		l := NewList()
		assert.Equal(t, 0, l.Len())
		assert.False(t, l.Contain("1"))
	})
	t.Run("string", func(t *testing.T) {
		assert.Equal(t, `[]`, NewList().String())
		assert.Equal(t, `[ "1" ]`, NewList("1").String())
		assert.Equal(t, `[ "1", "2" ]`, NewList("1", "2").String())
		assert.Equal(t, `[ "1", "", "2" ]`, NewList("1", "", "2").String())
	})
	t.Run("contain", func(t *testing.T) {
		l := NewList("1", "", "2")
		assert.True(t, l.Contain("1"))
		assert.True(t, l.Contain(""))
		assert.True(t, l.Contain("2"))
		assert.False(t, l.Contain("3"))
	})
	t.Run("clear", func(t *testing.T) {
		l := NewList("1", "2", "3")
		assert.Equal(t, 3, l.Len())
		l.Clear()
		assert.Equal(t, 0, l.Len())
		assert.Equal(t, `[]`, l.String())
	})
	t.Run("loop", func(t *testing.T) {
		t.Run("left", func(t *testing.T) {
			l := NewList("1", "2", "3")
			c1 := make([]string, l.Len())
			c2 := make([]string, l.Len())
			var j int
			l.Loop(Left, func(i int, v string) bool {
				if i == 2 {
					return false
				}
				c1[i] = v
				c2[j] = v
				j++
				return true
			})
			assert.Equal(t, []string{"1", "2", ""}, c1)
			assert.Equal(t, []string{"1", "2", ""}, c2)
		})
		t.Run("right", func(t *testing.T) {
			l := NewList("1", "2", "3")
			c1 := make([]string, l.Len(), l.Len())
			c2 := make([]string, l.Len(), l.Len())
			var j int
			l.Loop(Right, func(i int, v string) bool {
				if i == 0 {
					return false
				}
				c1[i] = v
				c2[j] = v
				j++
				return true
			})
			assert.Equal(t, []string{"", "2", "3"}, c1)
			assert.Equal(t, []string{"3", "2", ""}, c2)
		})
	})
	t.Run("first", func(t *testing.T) {
		l := NewList()

		v, ok := l.First()
		assert.False(t, ok)
		assert.Equal(t, "", v)

		l.Push(Left, "1")
		v, ok = l.First()
		assert.True(t, ok)
		assert.Equal(t, "1", v)

		l.Push(Left, "2")
		v, ok = l.First()
		assert.True(t, ok)
		assert.Equal(t, "2", v)
	})
	t.Run("last", func(t *testing.T) {
		l := NewList()

		v, ok := l.Last()
		assert.False(t, ok)
		assert.Equal(t, "", v)

		l.Push(Right, "1")
		v, ok = l.Last()
		assert.True(t, ok)
		assert.Equal(t, "1", v)

		l.Push(Right, "2")
		v, ok = l.Last()
		assert.True(t, ok)
		assert.Equal(t, "2", v)
	})
	t.Run("index of", func(t *testing.T) {
		l := NewList("3", "1", "", "1", "3")

		idx, ok := l.IndexOf(Left, "2")
		assert.False(t, ok)
		assert.Equal(t, 0, idx)

		idx, ok = l.IndexOf(Right, "2")
		assert.False(t, ok)
		assert.Equal(t, 0, idx)

		idx, ok = l.IndexOf(Left, "1")
		assert.True(t, ok)
		assert.Equal(t, 1, idx)

		idx, ok = l.IndexOf(Right, "1")
		assert.True(t, ok)
		assert.Equal(t, 3, idx)
	})
	t.Run("push", func(t *testing.T) {
		l := NewList()
		assert.Equal(t, 0, l.Len())
		assert.Equal(t, `[]`, l.String())

		l.Push(Right, "1", "2", "3")
		assert.Equal(t, 3, l.Len())
		assert.Equal(t, `[ "1", "2", "3" ]`, l.String())

		l.Push(Right, "4")
		assert.Equal(t, 4, l.Len())
		assert.Equal(t, `[ "1", "2", "3", "4" ]`, l.String())

		l.Push(Left, "5", "6")
		assert.Equal(t, 6, l.Len())
		assert.Equal(t, `[ "6", "5", "1", "2", "3", "4" ]`, l.String())
	})
	t.Run("pop", func(t *testing.T) {
		l := NewList("1", "2", "3")
		assert.Equal(t, 3, l.Len())
		assert.Equal(t, `[ "1", "2", "3" ]`, l.String())

		v, ok := l.Pop(Left)
		assert.True(t, ok)
		assert.Equal(t, "1", v)
		assert.Equal(t, 2, l.Len())
		assert.Equal(t, `[ "2", "3" ]`, l.String())

		v, ok = l.Pop(Right)
		assert.True(t, ok)
		assert.Equal(t, "3", v)
		assert.Equal(t, 1, l.Len())
		assert.Equal(t, `[ "2" ]`, l.String())

		v, ok = l.Pop(Left)
		assert.True(t, ok)
		assert.Equal(t, "2", v)
		assert.Equal(t, 0, l.Len())
		assert.Equal(t, `[]`, l.String())

		v, ok = l.Pop(Right)
		assert.False(t, ok)
		assert.Equal(t, "", v)
		assert.Equal(t, 0, l.Len())
		assert.Equal(t, `[]`, l.String())
	})
	t.Run("replace", func(t *testing.T) {
		l := NewList("1", "2", "3")

		v, err := l.Replace(1, "4")
		assert.NoError(t, err)
		assert.Equal(t, "2", v)

		v, err = l.Replace(10, "10")
		assert.Error(t, err)
		assert.True(t, errors.Is(err, ErrInvalidIndex))
		assert.Equal(t, "", v)
	})
}
