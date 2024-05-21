package passport

import (
	"cmp"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type SliceValue[T cmp.Ordered] struct {
	err        error
	key        string
	val        []T
	mark       bool
	localizer  *i18n.Localizer
	config     *i18n.LocalizeConfig
	messageMap map[string]*i18n.Message
}

func Slice[T cmp.Ordered](k string, v []T) *SliceValue[T] {
	return &SliceValue[T]{
		key: k,
		val: v,
	}
}

func (c *SliceValue[T]) setLocalizer(localizer *i18n.Localizer) {
	c.localizer = localizer
}

func (c *SliceValue[T]) validate(messageId string, ok bool, val any) *SliceValue[T] {
	if c.mark || ok {
		return c
	}
	c.mark = true
	messageId = "SliceValue." + messageId
	c.config = &i18n.LocalizeConfig{
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

	if c.localizer == nil {
		c.localizer = i18n.NewLocalizer(_bundle, defaultLanguages...)
	}

	if len(c.messageMap) > 0 {
		if message := c.messageMap[c.config.MessageID]; message != nil {
			c.config.DefaultMessage = message
			c.config.MessageID, c.config.DefaultMessage.ID = "!undefined", "!undefined"
			str, _ := c.localizer.Localize(c.config)
			c.err = errors.New(str)
			return c.err
		}
	}

	str, err := c.localizer.Localize(c.config)
	if err != nil {
		c.err = err
		return c.err
	}
	c.err = errors.New(str)
	return c.err
}

// Required the slice cannot be empty
func (c *SliceValue[T]) Required() *SliceValue[T] {
	return c.validate("Required", len(c.val) > 0, nil)
}

// Eq check the slice length is equal to v
func (c *SliceValue[T]) Eq(v int) *SliceValue[T] {
	return c.validate("Eq", len(c.val) == v, v)
}

// Gt check the slice length is greater than v
func (c *SliceValue[T]) Gt(v int) *SliceValue[T] {
	return c.validate("Gt", len(c.val) > v, v)
}

// Gte check the slice length is greater or equal than v
func (c *SliceValue[T]) Gte(v int) *SliceValue[T] {
	return c.validate("Gte", len(c.val) >= v, v)
}

// Lt check the slice length is less than v
func (c *SliceValue[T]) Lt(v int) *SliceValue[T] {
	return c.validate("Lt", len(c.val) < v, v)
}

// Lte check the slice length is less or equal than v
func (c *SliceValue[T]) Lte(v int) *SliceValue[T] {
	return c.validate("Lte", len(c.val) <= v, v)
}

// Include checks whether the slice contains v
func (c *SliceValue[T]) Include(v T) *SliceValue[T] {
	return c.validate("Include", contains(c.val, v), v)
}

// Exclude check if the slice does not contain v
func (c *SliceValue[T]) Exclude(v T) *SliceValue[T] {
	return c.validate("Exclude", !contains(c.val, v), v)
}

func (c *SliceValue[T]) Customize(layout string, f func([]T) bool) *SliceValue[T] {
	const funcName = "Customize"
	return c.validate(funcName, f(c.val), nil).Message(funcName, layout)
}

// Message customizing error messages
func (c *SliceValue[T]) Message(funcName string, layout string) *SliceValue[T] {
	if c.messageMap == nil {
		c.messageMap = make(map[string]*i18n.Message)
	}
	id := "SliceValue." + funcName
	c.messageMap[id] = &i18n.Message{
		ID:    id,
		Other: layout,
	}
	return c
}
