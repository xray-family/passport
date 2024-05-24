package passport

import "github.com/nicksnyder/go-i18n/v2/i18n"

type config struct {
	loc           *i18n.Localizer
	AutoTranslate bool
}

type Option func(c *config)

// WithLang set languages
func WithLang(langs ...string) Option {
	return func(c *config) {
		if len(langs) > 0 {
			c.loc = i18n.NewLocalizer(_bundle, langs...)
		}
	}
}

// WithAutoTranslate automatic translation of key names
func WithAutoTranslate() Option {
	return func(c *config) {
		c.AutoTranslate = true
	}
}

func withInit() Option {
	return func(c *config) {
		if c.loc == nil {
			c.loc = GetLocalizer()
		}
	}
}
