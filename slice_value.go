package passport

import (
	"cmp"
	"fmt"
)

type SlicedValue[T cmp.Ordered] struct {
	err error
	key string
	val []T
}

func Slice[T cmp.Ordered](k string, v []T) *SlicedValue[T] {
	return &SlicedValue[T]{
		key: k,
		val: v,
	}
}

func (c *SlicedValue[T]) validate(ok bool, layout string, args ...any) *SlicedValue[T] {
	if c.err != nil || ok {
		return c
	}
	c.err = fmt.Errorf(layout, args...)
	return c
}

// Err get error
func (c *SlicedValue[T]) Err() error {
	return c.err
}

// Required the slice cannot be empty
func (c *SlicedValue[T]) Required() *SlicedValue[T] {
	return c.validate(len(c.val) > 0, "%s is required", c.key)
}

// Eq check the slice length is equal to v
func (c *SlicedValue[T]) Eq(v int) *SlicedValue[T] {
	return c.validate(len(c.val) == v, "size of %s should equal %v", c.key, v)
}

// Gt check the slice length is greater than v
func (c *SlicedValue[T]) Gt(v int) *SlicedValue[T] {
	return c.validate(len(c.val) > v, "size of %s should great than %v", c.key, v)
}

// Gte check the slice length is greater or equal than v
func (c *SlicedValue[T]) Gte(v int) *SlicedValue[T] {
	return c.validate(len(c.val) >= v, "size of %s should great or equal than %v", c.key, v)
}

// Lt check the slice length is less than v
func (c *SlicedValue[T]) Lt(v int) *SlicedValue[T] {
	return c.validate(len(c.val) < v, "size of %s should less than %v", c.key, v)
}

// Lte check the slice length is less or equal than v
func (c *SlicedValue[T]) Lte(v int) *SlicedValue[T] {
	return c.validate(len(c.val) <= v, "size of %s should less or equal than %v", c.key, v)
}

// Include checks whether the slice contains v
func (c *SlicedValue[T]) Include(v T) *SlicedValue[T] {
	return c.validate(contains(c.val, v), "%s should contains %v", c.key, v)
}

// Exclude check if the slice does not contain v
func (c *SlicedValue[T]) Exclude(v T) *SlicedValue[T] {
	return c.validate(!contains(c.val, v), "%s should not contains %v", c.key, v)
}
