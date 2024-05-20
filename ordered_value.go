package passport

import (
	"cmp"
	"fmt"
)

type OrderedValue[T cmp.Ordered] struct {
	err error
	key string
	val T
}

func Ordered[T cmp.Ordered](k string, v T) *OrderedValue[T] {
	return &OrderedValue[T]{
		key: k,
		val: v,
	}
}

func (c *OrderedValue[T]) validate(ok bool, layout string, args ...any) *OrderedValue[T] {
	if c.err != nil || ok {
		return c
	}
	c.err = fmt.Errorf(layout, args...)
	return c
}

// Err get error
func (c *OrderedValue[T]) Err() error {
	return c.err
}

// Required the ordered value cannot be empty
func (c *OrderedValue[T]) Required() *OrderedValue[T] {
	return c.validate(!isZero(c.val), "%s is required", c.key)
}

// Gt check the ordered value is greater than v
func (c *OrderedValue[T]) Gt(v T) *OrderedValue[T] {
	return c.validate(c.val > v, "%s should great than %v", c.key, v)
}

// Gte check the ordered value is greater or equal than v
func (c *OrderedValue[T]) Gte(v T) *OrderedValue[T] {
	return c.validate(c.val >= v, "%s should great or equal than %v", c.key, v)
}

// Lt check the ordered value is less than v
func (c *OrderedValue[T]) Lt(v T) *OrderedValue[T] {
	return c.validate(c.val < v, "%s should less than %v", c.key, v)
}

// Lte check the ordered value is less or equal than v
func (c *OrderedValue[T]) Lte(v T) *OrderedValue[T] {
	return c.validate(c.val <= v, "%s should less or equal than %v", c.key, v)
}

// IncludeBy check if args contains the ordered value.
func (c *OrderedValue[T]) IncludeBy(args ...T) *OrderedValue[T] {
	return c.validate(contains(args, c.val), "%s should be one of %v", c.key, args)
}

// ExcludeBy checks if args does not contain the ordered value.
func (c *OrderedValue[T]) ExcludeBy(args ...T) *OrderedValue[T] {
	return c.validate(!contains(args, c.val), "%s should not be one of %v", c.key, args)
}
