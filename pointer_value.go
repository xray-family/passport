package passport

import (
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strconv"
)

type PointerValue[T any] struct {
	err     error
	key     string
	val     *T
	mark    bool
	conf    *config
	locConf *i18n.LocalizeConfig
}

func Pointer[T any](k string, v *T) *PointerValue[T] {
	return &PointerValue[T]{
		key:  k,
		val:  v,
		conf: _conf,
	}
}

func (c *PointerValue[T]) setConf(conf *config) {
	c.conf = conf
}

func (c *PointerValue[T]) validate(messageId string, ok bool, args ...any) *PointerValue[T] {
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

func (c *PointerValue[T]) Err() error {
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

func (c *PointerValue[T]) Required() *PointerValue[T] {
	return c.validate("PointerValue.Required", c.val != nil, nil)
}

func (c *PointerValue[T]) Customize(messageId string, f func(*T) bool) *PointerValue[T] {
	return c.validate(messageId, f(c.val), nil)
}
