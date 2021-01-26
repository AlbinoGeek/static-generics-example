package main

import (
	"errors"

	"github.com/cheekybits/genny/generic"
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

func (c *ItemList) Get(i int) (val Item, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *ItemList) Set(idx int, val Item) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}
