package passport

import (
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type PointerValue[T any] struct {
	err       error
	key       string
	val       *T
	mark      bool
	localizer *i18n.Localizer
	config    *i18n.LocalizeConfig
}

func Pointer[T any](k string, v *T) *PointerValue[T] {
	return &PointerValue[T]{
		key:       k,
		val:       v,
		localizer: GetLocalizer(),
	}
}

func (c *PointerValue[T]) setLocalizer(localizer *i18n.Localizer) {
	c.localizer = localizer
}

func (c *PointerValue[T]) validate(messageId string, ok bool, val any) *PointerValue[T] {
	if c.mark || ok {
		return c
	}
	c.mark = true
	c.config = &i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: map[string]any{"Key": c.key, "Value": val},
	}
	return c
}

func (c *PointerValue[T]) Err() error {
	if !c.mark {
		return nil
	}
	if c.err != nil {
		return c.err
	}

	str, err := c.localizer.Localize(c.config)
	if err != nil {
		c.err = err
		return c.err
	}

	c.err = errors.New(str)
	return c.err
}

func (c *PointerValue[T]) Required() *PointerValue[T] {
	return c.validate("PointerValue.Required", c.val != nil, nil)
}

func (c *PointerValue[T]) Customize(messageId string, f func(*T) bool) *PointerValue[T] {
	return c.validate(messageId, f(c.val), nil)
}
