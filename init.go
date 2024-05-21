package passport

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	defaultLanguages = []string{"en_US"}
	_bundle          *i18n.Bundle
)

func SetLang(langs ...string) {
	defaultLanguages = langs
}

func init() {
	_bundle = i18n.NewBundle(language.English)
	_bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_bundle.MustLoadMessageFile(`active.en-US.toml`)
	_bundle.MustLoadMessageFile(`active.zh-CN.toml`)
}
