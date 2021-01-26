package main

import (
	"testing"

	"github.com/cheekybits/genny/generic"
	"gopkg.in/go-playground/assert.v1"
)

type Item generic.Type

type ItemList struct {
	s []Item
}

func NewItemList() *ItemList {
	return &ItemList{s: []Item{}}
}

func (c *ItemList) Add(val Item) {
	c.s = append(c.s, val)
}

func (c *ItemList) Remove(val Item) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *ItemList) Get(i int) (v Item) {
	if i < len(c.s) {
		v = c.s[i]
	}

	return
}

func (c *ItemList) Set(i int, v Item) {
	c.s[i] = v
}

func TestItemList(t *testing.T) {
	l := NewItemList()

	// TODO: More meaningful test
	var v Item

	l.Add(v)
	assert.Equal(t, v, l.Get(0))

	l.Set(0, v)
	assert.Equal(t, v, l.Get(0))

	l.Remove(v)
	assert.Equal(t, "", l.Get(0))
}
