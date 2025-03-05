package passport

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type config struct {
	loc *i18n.Localizer
}

type Option func(c *config)

// withLang set languages
func withLang(r *http.Request) Option {
	return func(c *config) {
		if r != nil {
			lang := r.FormValue("lang")
			accept := r.Header.Get("Accept-Language")
			c.loc = i18n.NewLocalizer(GetBundle(), lang, accept)
		}
	}
}

func withInit() Option {
	return func(c *config) {
		if c.loc == nil {
			c.loc = GetLocalizer()
		}
	}
}
