package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemList(t *testing.T) {
	l := NewItemList()

	// TODO: More meaningful test
	var v Item

	l.Append(v)

	x, e := l.Get(0)
	assert.Equal(t, v, x)
	assert.NoError(t, e)

	l.Set(0, v)
	x, e = l.Get(0)
	assert.Equal(t, v, x)
	assert.NoError(t, e)

	l.Remove(v)
	x, e = l.Get(0)
	assert.Equal(t, v, x)
	assert.Error(t, e)
}
