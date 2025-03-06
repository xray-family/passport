package validator

import (
	"embed"
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	//go:embed asset/active.en-US.json
	enFS embed.FS

	//go:embed asset/active.zh-CN.json
	cnFS embed.FS

	_bundle    *i18n.Bundle
	_localizer *i18n.Localizer
	_conf      *config

	Chinese = language.Make("zh-CN")
	English = language.Make("en-US")
)

func init() {
	SetLang(English, English.String())
}

func SetLang(tag language.Tag, langs ...string) {
	_bundle = i18n.NewBundle(tag)
	_bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	_, _ = _bundle.LoadMessageFileFS(enFS, `asset/active.en-US.json`)
	_, _ = _bundle.LoadMessageFileFS(cnFS, `asset/active.zh-CN.json`)
	_localizer = i18n.NewLocalizer(_bundle, langs...)
	_conf = new(config)
	withInit()(_conf)
}

func GetBundle() *i18n.Bundle {
	return _bundle
}

func GetLocalizer() *i18n.Localizer {
	return _localizer
}
