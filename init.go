package passport

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	_lang      []string
	_bundle    *i18n.Bundle
	_localizer *i18n.Localizer
	_conf      *config

	Chinese = language.Make("zh-CN")
	English = language.Make("en-US")
)

func init() {
	SetLang(English, English.String())

	_conf = new(config)
	withInit()(_conf)
}

func SetLang(tag language.Tag, langs ...string) {
	_lang = langs
	_bundle = i18n.NewBundle(tag)
	_bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_bundle.MustLoadMessageFile(`active.en-US.toml`)
	_bundle.MustLoadMessageFile(`active.zh-CN.toml`)
	_localizer = i18n.NewLocalizer(_bundle, langs...)
}

func GetBundle() *i18n.Bundle {
	return _bundle
}

func GetLocalizer() *i18n.Localizer {
	return _localizer
}
