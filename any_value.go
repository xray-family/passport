package passport

import (
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AnyValue struct {
	err       error
	key       string
	val       any
	mark      bool
	localizer *i18n.Localizer
	config    *i18n.LocalizeConfig
}

func Any(k string, v any) *AnyValue {
	return &AnyValue{
		key: k,
		val: v,
	}
}

func (c *AnyValue) validate(messageId string, ok bool, val any) *AnyValue {
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

// Err get error
func (c *AnyValue) Err() error {
	if !c.mark {
		return nil
	}
	if c.err != nil {
		return c.err
	}

	if c.localizer == nil {
		c.localizer = i18n.NewLocalizer(_bundle, _lang...)
	}

	str, err := c.localizer.Localize(c.config)
	if err != nil {
		c.err = err
		return c.err
	}
	c.err = errors.New(str)
	return c.err
}

func (c *AnyValue) setLocalizer(localizer *i18n.Localizer) {
	c.localizer = localizer
}

func (c *AnyValue) Customize(messageId string, f func(any) bool) *AnyValue {
	return c.validate(messageId, f(c.val), nil)
}
