package passport

import (
	"cmp"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type OrderedValue[T cmp.Ordered] struct {
	err        error
	key        string
	val        T
	mark       bool
	localizer  *i18n.Localizer
	config     *i18n.LocalizeConfig
	messageMap map[string]*i18n.Message
}

func Ordered[T cmp.Ordered](k string, v T) *OrderedValue[T] {
	return &OrderedValue[T]{
		key: k,
		val: v,
	}
}

func (c *OrderedValue[T]) validate(messageId string, ok bool, val any) *OrderedValue[T] {
	if c.mark || ok {
		return c
	}
	c.mark = true
	messageId = "OrderedValue." + messageId
	c.config = &i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: map[string]any{"Key": c.key, "Value": val},
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

func (c *OrderedValue[T]) setLocalizer(localizer *i18n.Localizer) {
	c.localizer = localizer
}

// Required the ordered value cannot be empty
func (c *OrderedValue[T]) Required() *OrderedValue[T] {
	return c.validate("Required", !isZero(c.val), nil)
}

// Gt check the ordered value is greater than v
func (c *OrderedValue[T]) Gt(v T) *OrderedValue[T] {
	return c.validate("Gt", c.val > v, v)
}

// Gte check the ordered value is greater or equal than v
func (c *OrderedValue[T]) Gte(v T) *OrderedValue[T] {
	return c.validate("Gte", c.val >= v, v)
}

// Lt check the ordered value is less than v
func (c *OrderedValue[T]) Lt(v T) *OrderedValue[T] {
	return c.validate("Lt", c.val < v, v)
}

// Lte check the ordered value is less or equal than v
func (c *OrderedValue[T]) Lte(v T) *OrderedValue[T] {
	return c.validate("Lte", c.val <= v, v)
}

// IncludeBy check if args contains the ordered value.
func (c *OrderedValue[T]) IncludeBy(args ...T) *OrderedValue[T] {
	return c.validate("IncludeBy", contains(args, c.val), args)
}

// ExcludeBy checks if args does not contain the ordered value.
func (c *OrderedValue[T]) ExcludeBy(args ...T) *OrderedValue[T] {
	return c.validate("ExcludeBy", !contains(args, c.val), args)
}

// Customize customized data validation
// @layout error message
// @f check function
func (c *OrderedValue[T]) Customize(layout string, f func(T) bool) *OrderedValue[T] {
	const funcName = "Customize"
	return c.validate(funcName, f(c.val), nil).Message(funcName, layout)
}

// Message customizing error messages
func (c *OrderedValue[T]) Message(funcName string, layout string) *OrderedValue[T] {
	if c.messageMap == nil {
		c.messageMap = make(map[string]*i18n.Message)
	}
	id := "OrderedValue." + funcName
	c.messageMap[id] = &i18n.Message{
		ID:    id,
		Other: layout,
	}
	return c
}
