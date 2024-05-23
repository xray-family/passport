package passport

import (
	"cmp"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SliceValue[T cmp.Ordered] struct {
	err     error
	key     string
	val     []T
	mark    bool
	conf    *config
	locConf *i18n.LocalizeConfig
}

func Slice[T cmp.Ordered](k string, v []T) *SliceValue[T] {
	return &SliceValue[T]{
		key:  k,
		val:  v,
		conf: _conf,
	}
}

func (c *SliceValue[T]) setConf(conf *config) {
	c.conf = conf
}

func (c *SliceValue[T]) validate(messageId string, ok bool, val any) *SliceValue[T] {
	if c.mark || ok {
		return c
	}
	c.mark = true
	c.locConf = &i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: map[string]any{"Key": c.key, "Value": val},
	}
	return c
}

// Err get error
func (c *SliceValue[T]) Err() error {
	if !c.mark {
		return nil
	}
	if c.err != nil {
		return c.err
	}
	if c.conf.AutoTranslate {
		if str, err := c.conf.loc.Localize(&i18n.LocalizeConfig{MessageID: c.key}); err == nil {
			td := c.locConf.TemplateData.(map[string]any)
			td["Key"] = str
		}
	}
	str, err := c.conf.loc.Localize(c.locConf)
	if err != nil {
		c.err = err
		return c.err
	}
	c.err = errors.New(str)
	return c.err
}

// Required the slice cannot be empty
func (c *SliceValue[T]) Required() *SliceValue[T] {
	return c.validate("SliceValue.Required", len(c.val) > 0, nil)
}

// Eq check the slice length is equal to v
func (c *SliceValue[T]) Eq(v int) *SliceValue[T] {
	return c.validate("SliceValue.Eq", len(c.val) == v, v)
}

// Gt check the slice length is greater than v
func (c *SliceValue[T]) Gt(v int) *SliceValue[T] {
	return c.validate("SliceValue.Gt", len(c.val) > v, v)
}

// Gte check the slice length is greater or equal than v
func (c *SliceValue[T]) Gte(v int) *SliceValue[T] {
	return c.validate("SliceValue.Gte", len(c.val) >= v, v)
}

// Lt check the slice length is less than v
func (c *SliceValue[T]) Lt(v int) *SliceValue[T] {
	return c.validate("SliceValue.Lt", len(c.val) < v, v)
}

// Lte check the slice length is less or equal than v
func (c *SliceValue[T]) Lte(v int) *SliceValue[T] {
	return c.validate("SliceValue.Lte", len(c.val) <= v, v)
}

// Contains checks whether the slice contains v
func (c *SliceValue[T]) Contains(v T) *SliceValue[T] {
	return c.validate("SliceValue.Contains", contains(c.val, v), v)
}

func (c *SliceValue[T]) Customize(messageId string, f func([]T) bool) *SliceValue[T] {
	return c.validate(messageId, f(c.val), nil)
}
