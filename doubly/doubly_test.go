package doubly

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	l := New[int]()
	assert.True(t, l.IsEmpty())

	l.AddFirst(1)
	assert.False(t, l.IsEmpty())

	l.RemoveFirst()
	assert.True(t, l.IsEmpty())

	l.AddLast(1)
	assert.False(t, l.IsEmpty())

	l.RemoveLast()
	assert.True(t, l.IsEmpty())
}

func TestFirst(t *testing.T) {
	l := New[int]()

	f, ok := l.First()
	assert.False(t, ok)
	assert.Equal(t, 0, f)

	l.AddFirst(3)
	l.AddFirst(2)
	l.AddFirst(1)

	f, ok = l.First()
	assert.True(t, ok)
	assert.Equal(t, 1, f)
}

func TestLast(t *testing.T) {
	list := New[int]()

	l, ok := list.Last()
	assert.False(t, ok)
	assert.Equal(t, 0, l)

	list.AddFirst(3)
	list.AddFirst(2)
	list.AddFirst(1)

	l, ok = list.Last()
	assert.True(t, ok)
	assert.Equal(t, 3, l)
}

func TestAddBetween(t *testing.T) {
	list := New[int]()

	list.AddBetween(10, list.header, list.trailer)

	f, ok := list.First()
	assert.True(t, ok)
	assert.Equal(t, 10, f)
}

func TestAddFirst(t *testing.T) {
	l := New[int]()

	l.AddFirst(1)
	l.AddFirst(5)

	f, ok := l.First()
	assert.True(t, ok)
	assert.Equal(t, 5, f)
}

func TestAddLast(t *testing.T) {
	list := New[int]()

	list.AddLast(1)
	list.AddLast(5)

	l, ok := list.Last()
	assert.True(t, ok)
	assert.Equal(t, 5, l)
}

func TestRemove(t *testing.T) {
	list := New[int]()

	list.AddFirst(10)
	list.AddLast(11)

	list.Remove(list.header.Next)

	f, ok := list.First()
	assert.True(t, ok)
	assert.Equal(t, 11, f)
}

func TestRemoveFirst(t *testing.T) {
	list := New[int]()

	f, ok := list.RemoveFirst()
	assert.False(t, ok)
	assert.Equal(t, 0, f)

	list.AddFirst(10)
	list.AddFirst(11)
	list.AddFirst(12)

	f, ok = list.RemoveFirst()
	assert.True(t, ok)
	assert.Equal(t, 12, f)

	f, ok = list.First()
	assert.True(t, ok)
	assert.Equal(t, 11, f)
}

func TestRemoveLast(t *testing.T) {
	list := New[int]()

	l, ok := list.RemoveLast()
	assert.False(t, ok)
	assert.Equal(t, 0, l)

	list.AddFirst(10)
	list.AddFirst(11)
	list.AddFirst(12)

	l, ok = list.RemoveLast()
	assert.True(t, ok)
	assert.Equal(t, 10, l)

	l, ok = list.Last()
	assert.True(t, ok)
	assert.Equal(t, 11, l)
}

func TestListString(t *testing.T) {
	list := New[int]()

	s := list.String()
	assert.Equal(t, "[ ]", s)

	list.AddFirst(1)
	s = list.String()
	assert.Equal(t, "[ 1 ]", s)

	list.AddFirst(2)
	list.AddFirst(3)
	s = list.String()
	assert.Equal(t, "[ 3 2 1 ]", s)
}

func TestToSlice(t *testing.T) {
	list := New[int]()

	slice := list.ToSlice()
	s := fmt.Sprint(slice)
	assert.Equal(t, "[]", s)

	list.AddFirst(1)
	slice = list.ToSlice()
	s = fmt.Sprint(slice)
	assert.Equal(t, "[1]", s)

	list.AddFirst(2)
	slice = list.ToSlice()
	s = fmt.Sprint(slice)
	assert.Equal(t, "[2 1]", s)
}
