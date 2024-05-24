package passport

import (
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strconv"
)

type AnyValue[T any] struct {
	err     error
	key     string
	val     T
	mark    bool
	conf    *config
	locConf *i18n.LocalizeConfig
}

func Any[T any](k string, v T) *AnyValue[T] {
	return &AnyValue[T]{
		key:  k,
		val:  v,
		conf: _conf,
	}
}

func (c *AnyValue[T]) validate(messageId string, ok bool, args ...any) *AnyValue[T] {
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
func (c *AnyValue[T]) Err() error {
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

func (c *AnyValue[T]) setConf(conf *config) {
	c.conf = conf
}

func (c *AnyValue[T]) Customize(messageId string, f func(T) bool) *AnyValue[T] {
	return c.validate(messageId, f(c.val), nil)
}
