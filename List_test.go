// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringList(t *testing.T) {
	l := NewStringList()

	// TODO: More meaningful test
	var v string

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

func TestIntList(t *testing.T) {
	l := NewIntList()

	// TODO: More meaningful test
	var v int

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

func TestInt32List(t *testing.T) {
	l := NewInt32List()

	// TODO: More meaningful test
	var v int32

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

func TestInt64List(t *testing.T) {
	l := NewInt64List()

	// TODO: More meaningful test
	var v int64

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

func TestBoolList(t *testing.T) {
	l := NewBoolList()

	// TODO: More meaningful test
	var v bool

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
