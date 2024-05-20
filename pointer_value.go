package passport

import (
	"fmt"
)

type PointerValue[T any] struct {
	err   error
	key   string
	value *T
}

func Pointer[T any](k string, v *T) *PointerValue[T] {
	return &PointerValue[T]{
		key:   k,
		value: v,
	}
}

func (c *PointerValue[T]) validate(ok bool, layout string, args ...any) *PointerValue[T] {
	if c.err != nil || ok {
		return c
	}
	c.err = fmt.Errorf(layout, args...)
	return c
}

func (c *PointerValue[T]) Err() error {
	return c.err
}

func (c *PointerValue[T]) Required() *PointerValue[T] {
	return c.validate(c.value != nil, "%s is required", c.key)
}
