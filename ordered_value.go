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

func (c *OrderedValue[T]) Err() error {
	return c.err
}

func (c *OrderedValue[T]) Required() *OrderedValue[T] {
	return c.validate(!isZero(c.val), "%s is required", c.key)
}

func (c *OrderedValue[T]) Gt(v T) *OrderedValue[T] {
	return c.validate(c.val > v, "%s should great than %v", c.key, v)
}

func (c *OrderedValue[T]) Gte(v T) *OrderedValue[T] {
	return c.validate(c.val >= v, "%s should great or equal than %v", c.key, v)
}

func (c *OrderedValue[T]) Lt(v T) *OrderedValue[T] {
	return c.validate(c.val < v, "%s should less than %v", c.key, v)
}

func (c *OrderedValue[T]) Lte(v T) *OrderedValue[T] {
	return c.validate(c.val <= v, "%s should less or equal than %v", c.key, v)
}

func (c *OrderedValue[T]) IncludeBy(args ...T) *OrderedValue[T] {
	return c.validate(contains(args, c.val), "%s should be one of %v", c.key, args)
}

func (c *OrderedValue[T]) ExcludeBy(args ...T) *OrderedValue[T] {
	return c.validate(!contains(args, c.val), "%s should not be one of %v", c.key, args)
}
