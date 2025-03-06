package validator

import (
	"cmp"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strconv"
)

type OrderedValue[T cmp.Ordered] struct {
	err     error
	key     string
	val     T
	mark    bool
	conf    *config
	locConf *i18n.LocalizeConfig
}

func Ordered[T cmp.Ordered](k string, v T) *OrderedValue[T] {
	return &OrderedValue[T]{
		key:  k,
		val:  v,
		conf: _conf,
	}
}

func (c *OrderedValue[T]) validate(messageId string, ok bool, args ...any) *OrderedValue[T] {
	if c.mark || ok {
		return c
	}
	c.mark = true
	td := map[string]any{"Key": c.key}
	for i, v := range args {
		key := "Arg" + strconv.Itoa(i)
		td[key] = v
	}
	c.locConf = &i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: td,
	}
	return c
}

// Err get error
func (c *OrderedValue[T]) Err() error {
	if !c.mark {
		return nil
	}
	if c.err != nil {
		return c.err
	}
	str, err := c.conf.loc.Localize(c.locConf)
	if err != nil {
		c.err = err
		return c.err
	}
	c.err = errors.New(str)
	return c.err
}

func (c *OrderedValue[T]) setConf(conf *config) {
	c.conf = conf
}

// Required the ordered value cannot be empty
func (c *OrderedValue[T]) Required() *OrderedValue[T] {
	return c.validate("OrderedValue.Required", !isZero(c.val), nil)
}

// Gt check the ordered value is greater than v
func (c *OrderedValue[T]) Gt(v T) *OrderedValue[T] {
	return c.validate("OrderedValue.Gt", c.val > v, v)
}

// Gte check the ordered value is greater or equal than v
func (c *OrderedValue[T]) Gte(v T) *OrderedValue[T] {
	return c.validate("OrderedValue.Gte", c.val >= v, v)
}

// Lt check the ordered value is less than v
func (c *OrderedValue[T]) Lt(v T) *OrderedValue[T] {
	return c.validate("OrderedValue.Lt", c.val < v, v)
}

// Lte check the ordered value is less or equal than v
func (c *OrderedValue[T]) Lte(v T) *OrderedValue[T] {
	return c.validate("OrderedValue.Lte", c.val <= v, v)
}

// Between check that the range of values of the ordered value satisfies a <= x <b
func (c *OrderedValue[T]) Between(a, b T) *OrderedValue[T] {
	return c.validate("OrderedValue.Between", c.val >= a && c.val < b, a, b)
}

// In check if args contains the ordered value.
func (c *OrderedValue[T]) In(args ...T) *OrderedValue[T] {
	return c.validate("OrderedValue.In", contains(args, c.val), args)
}

// Customize customized data validation
// @layout error message
// @f check function
func (c *OrderedValue[T]) Customize(messageId string, f func(T) bool) *OrderedValue[T] {
	return c.validate(messageId, f(c.val), nil)
}
