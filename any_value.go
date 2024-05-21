package passport

import (
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type AnyValue struct {
	err        error
	key        string
	val        any
	mark       bool
	langs      []string
	localizer  *i18n.Localizer
	config     *i18n.LocalizeConfig
	messageMap map[string]*i18n.Message
}

func Any(k string, v any) *AnyValue {
	return &AnyValue{
		langs: defaultLanguages,
		key:   k,
		val:   v,
	}
}

func (c *AnyValue) validate(messageId string, ok bool, val any) *AnyValue {
	if c.mark || ok {
		return c
	}
	c.mark = true
	messageId = "AnyValue." + messageId
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
		c.localizer = i18n.NewLocalizer(_bundle, c.langs...)
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

func (c *AnyValue) setLang(langs ...string) {
	c.localizer = i18n.NewLocalizer(_bundle, langs...)
	c.langs = langs
}

func (c *AnyValue) Customize(layout string, f func(any) bool) *AnyValue {
	const funcName = "Customize"
	return c.validate(funcName, f(c.val), nil).Message(funcName, layout)
}

// Message customizing error messages
func (c *AnyValue) Message(funcName string, layout string) *AnyValue {
	if c.messageMap == nil {
		c.messageMap = make(map[string]*i18n.Message)
	}
	id := "AnyValue." + funcName
	c.messageMap[id] = &i18n.Message{
		ID:    id,
		Other: layout,
	}
	return c
}
