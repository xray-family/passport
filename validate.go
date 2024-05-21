package passport

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type Value interface {
	setLocalizer(localizer *i18n.Localizer)
	Err() error
}

type Validator struct {
	localizer *i18n.Localizer
}

func NewValidator(langs ...string) *Validator {
	return &Validator{
		localizer: i18n.NewLocalizer(_bundle, langs...),
	}
}

func NewValidatorFromHTTP(header http.Header) *Validator {
	return &Validator{
		localizer: i18n.NewLocalizer(_bundle, header.Get("Accept-Language")),
	}
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

func Validate(values ...Value) error {
	for _, item := range values {
		if err := item.Err(); err != nil {
			return err
		}
	}
	return nil
}
