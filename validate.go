package passport

import "net/http"

type Value interface {
	setConf(conf *config)
	Err() error
}

type Validator struct {
	conf *config
}

func NewValidator(r *http.Request, options ...Option) *Validator {
	options = append(options, withLang(r), withInit())
	var conf = new(config)
	for _, f := range options {
		f(conf)
	}
	return &Validator{conf: conf}
}

func (c *Validator) Validate(values ...Value) error {
	for _, item := range values {
		item.setConf(c.conf)
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
