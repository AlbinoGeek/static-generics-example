// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package main

import "errors"

type StringList struct {
	s []string
}

func NewStringList() *StringList {
	return &StringList{s: []string{}}
}

func (c *StringList) Append(val string) {
	c.s = append(c.s, val)
}

func (c *StringList) Prepend(val string) {
	c.s = append([]string{val}, c.s...)
}

func (c *StringList) Remove(val string) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *StringList) Get(i int) (val string, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *StringList) Set(idx int, val string) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type IntList struct {
	s []int
}

func NewIntList() *IntList {
	return &IntList{s: []int{}}
}

func (c *IntList) Append(val int) {
	c.s = append(c.s, val)
}

func (c *IntList) Prepend(val int) {
	c.s = append([]int{val}, c.s...)
}

func (c *IntList) Remove(val int) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *IntList) Get(i int) (val int, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *IntList) Set(idx int, val int) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Int8List struct {
	s []int8
}

func NewInt8List() *Int8List {
	return &Int8List{s: []int8{}}
}

func (c *Int8List) Append(val int8) {
	c.s = append(c.s, val)
}

func (c *Int8List) Prepend(val int8) {
	c.s = append([]int8{val}, c.s...)
}

func (c *Int8List) Remove(val int8) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Int8List) Get(i int) (val int8, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Int8List) Set(idx int, val int8) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Int16List struct {
	s []int16
}

func NewInt16List() *Int16List {
	return &Int16List{s: []int16{}}
}

func (c *Int16List) Append(val int16) {
	c.s = append(c.s, val)
}

func (c *Int16List) Prepend(val int16) {
	c.s = append([]int16{val}, c.s...)
}

func (c *Int16List) Remove(val int16) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Int16List) Get(i int) (val int16, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Int16List) Set(idx int, val int16) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Int32List struct {
	s []int32
}

func NewInt32List() *Int32List {
	return &Int32List{s: []int32{}}
}

func (c *Int32List) Append(val int32) {
	c.s = append(c.s, val)
}

func (c *Int32List) Prepend(val int32) {
	c.s = append([]int32{val}, c.s...)
}

func (c *Int32List) Remove(val int32) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Int32List) Get(i int) (val int32, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Int32List) Set(idx int, val int32) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Int64List struct {
	s []int64
}

func NewInt64List() *Int64List {
	return &Int64List{s: []int64{}}
}

func (c *Int64List) Append(val int64) {
	c.s = append(c.s, val)
}

func (c *Int64List) Prepend(val int64) {
	c.s = append([]int64{val}, c.s...)
}

func (c *Int64List) Remove(val int64) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Int64List) Get(i int) (val int64, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Int64List) Set(idx int, val int64) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type BoolList struct {
	s []bool
}

func NewBoolList() *BoolList {
	return &BoolList{s: []bool{}}
}

func (c *BoolList) Append(val bool) {
	c.s = append(c.s, val)
}

func (c *BoolList) Prepend(val bool) {
	c.s = append([]bool{val}, c.s...)
}

func (c *BoolList) Remove(val bool) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *BoolList) Get(i int) (val bool, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *BoolList) Set(idx int, val bool) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type UintList struct {
	s []uint
}

func NewUintList() *UintList {
	return &UintList{s: []uint{}}
}

func (c *UintList) Append(val uint) {
	c.s = append(c.s, val)
}

func (c *UintList) Prepend(val uint) {
	c.s = append([]uint{val}, c.s...)
}

func (c *UintList) Remove(val uint) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *UintList) Get(i int) (val uint, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *UintList) Set(idx int, val uint) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Uint8List struct {
	s []uint8
}

func NewUint8List() *Uint8List {
	return &Uint8List{s: []uint8{}}
}

func (c *Uint8List) Append(val uint8) {
	c.s = append(c.s, val)
}

func (c *Uint8List) Prepend(val uint8) {
	c.s = append([]uint8{val}, c.s...)
}

func (c *Uint8List) Remove(val uint8) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Uint8List) Get(i int) (val uint8, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Uint8List) Set(idx int, val uint8) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Uint16List struct {
	s []uint16
}

func NewUint16List() *Uint16List {
	return &Uint16List{s: []uint16{}}
}

func (c *Uint16List) Append(val uint16) {
	c.s = append(c.s, val)
}

func (c *Uint16List) Prepend(val uint16) {
	c.s = append([]uint16{val}, c.s...)
}

func (c *Uint16List) Remove(val uint16) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Uint16List) Get(i int) (val uint16, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Uint16List) Set(idx int, val uint16) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Uint32List struct {
	s []uint32
}

func NewUint32List() *Uint32List {
	return &Uint32List{s: []uint32{}}
}

func (c *Uint32List) Append(val uint32) {
	c.s = append(c.s, val)
}

func (c *Uint32List) Prepend(val uint32) {
	c.s = append([]uint32{val}, c.s...)
}

func (c *Uint32List) Remove(val uint32) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Uint32List) Get(i int) (val uint32, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Uint32List) Set(idx int, val uint32) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Uint64List struct {
	s []uint64
}

func NewUint64List() *Uint64List {
	return &Uint64List{s: []uint64{}}
}

func (c *Uint64List) Append(val uint64) {
	c.s = append(c.s, val)
}

func (c *Uint64List) Prepend(val uint64) {
	c.s = append([]uint64{val}, c.s...)
}

func (c *Uint64List) Remove(val uint64) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Uint64List) Get(i int) (val uint64, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Uint64List) Set(idx int, val uint64) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Float32List struct {
	s []float32
}

func NewFloat32List() *Float32List {
	return &Float32List{s: []float32{}}
}

func (c *Float32List) Append(val float32) {
	c.s = append(c.s, val)
}

func (c *Float32List) Prepend(val float32) {
	c.s = append([]float32{val}, c.s...)
}

func (c *Float32List) Remove(val float32) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Float32List) Get(i int) (val float32, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Float32List) Set(idx int, val float32) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Float64List struct {
	s []float64
}

func NewFloat64List() *Float64List {
	return &Float64List{s: []float64{}}
}

func (c *Float64List) Append(val float64) {
	c.s = append(c.s, val)
}

func (c *Float64List) Prepend(val float64) {
	c.s = append([]float64{val}, c.s...)
}

func (c *Float64List) Remove(val float64) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Float64List) Get(i int) (val float64, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Float64List) Set(idx int, val float64) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Complex64List struct {
	s []complex64
}

func NewComplex64List() *Complex64List {
	return &Complex64List{s: []complex64{}}
}

func (c *Complex64List) Append(val complex64) {
	c.s = append(c.s, val)
}

func (c *Complex64List) Prepend(val complex64) {
	c.s = append([]complex64{val}, c.s...)
}

func (c *Complex64List) Remove(val complex64) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Complex64List) Get(i int) (val complex64, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Complex64List) Set(idx int, val complex64) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

type Complex128List struct {
	s []complex128
}

func NewComplex128List() *Complex128List {
	return &Complex128List{s: []complex128{}}
}

func (c *Complex128List) Append(val complex128) {
	c.s = append(c.s, val)
}

func (c *Complex128List) Prepend(val complex128) {
	c.s = append([]complex128{val}, c.s...)
}

func (c *Complex128List) Remove(val complex128) {
	for i, v := range c.s {
		if v == val {
			c.s = append(c.s[:i], c.s[i+1:]...)
			return
		}
	}
}

func (c *Complex128List) Get(i int) (val complex128, err error) {
	if i < len(c.s) {
		val = c.s[i]
	} else {
		err = errors.New("index out of bounds")
	}

	return
}

func (c *Complex128List) Set(idx int, val complex128) (err error) {
	if idx < len(c.s) {
		c.s[idx] = val
	} else {
		err = errors.New("index out of bounds")
	}

	return
}
