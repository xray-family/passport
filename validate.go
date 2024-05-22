package passport

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type Value interface {
	setLocalizer(localizer *i18n.Localizer)
	Err() error
}

type Validator struct {
	localizer *i18n.Localizer
}

func NewValidator(langs ...string) *Validator {
	var localizer = GetLocalizer()
	if len(langs) > 0 {
		localizer = i18n.NewLocalizer(_bundle, langs...)
	}
	return &Validator{localizer: localizer}
}

func (c *Validator) Validate(values ...Value) error {
	for _, item := range values {
		item.setLocalizer(c.localizer)
		if err := item.Err(); err != nil {
			return err
		}
	}
	return nil
}

// Localize translation into the given language
func (c *Validator) Localize(messageId string) string {
	s, _ := c.localizer.Localize(&i18n.LocalizeConfig{MessageID: messageId})
	return s
}

func Validate(values ...Value) error {
	for _, item := range values {
		if err := item.Err(); err != nil {
			return err
		}
	}
	return nil
}
