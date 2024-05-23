package passport

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"regexp"
	"strings"
)

var (
	reAlphabet        = regexp.MustCompile(`^[A-Za-z]+$`)
	reNumeric         = regexp.MustCompile(`^[0-9]+$`)
	reAlphabetNumeric = regexp.MustCompile(`^[A-Za-z0-9]+$`)
)

type StringValue[T ~string] struct {
	err     error
	key     string
	val     T
	mark    bool
	conf    *config
	locConf *i18n.LocalizeConfig
}

func String[T ~string](k string, v T) *StringValue[T] {
	return &StringValue[T]{
		key:  k,
		val:  T(strings.TrimSpace(string(v))),
		conf: _conf,
	}
}

func (c *StringValue[T]) setConf(conf *config) {
	c.conf = conf
}

func (c *StringValue[T]) validate(messageId string, ok bool, val any) *StringValue[T] {
	if c.mark || ok {
		return c
	}
	c.mark = true
	c.locConf = &i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: map[string]any{"Key": c.key, "Value": val},
	}
	return c
}

// Err get error
func (c *StringValue[T]) Err() error {
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

// Required the string cannot be empty
func (c *StringValue[T]) Required() *StringValue[T] {
	return c.validate("StringValue.Required", !isZero(c.val), nil)
}

// Eq check that the string length is equal to v
func (c *StringValue[T]) Eq(v int) *StringValue[T] {
	return c.validate("StringValue.Eq", len(c.val) == v, v)
}

// Gt check that the string length is greater than v
func (c *StringValue[T]) Gt(v int) *StringValue[T] {
	return c.validate("StringValue.Gt", len(c.val) > v, v)
}

// Gte check that the string length is greater or equal than v
func (c *StringValue[T]) Gte(v int) *StringValue[T] {
	return c.validate("StringValue.gte", len(c.val) >= v, v)
}

// Lt check that the string length is less than v
func (c *StringValue[T]) Lt(v int) *StringValue[T] {
	return c.validate("StringValue.Lt", len(c.val) < v, v)
}

// Lte check that the string length is less or equal than v
func (c *StringValue[T]) Lte(v int) *StringValue[T] {
	return c.validate("StringValue.Lte", len(c.val) <= v, v)
}

// In check if args contains the string.
func (c *StringValue[T]) In(args ...T) *StringValue[T] {
	return c.validate("StringValue.In", contains(args, c.val), args)
}

// MatchString verify that the string matches the regular expression re
func (c *StringValue[T]) MatchString(re string) *StringValue[T] {
	if c.mark {
		return c
	}
	r, err := regexp.Compile(re)
	if err != nil {
		return c.validate("StringValue.ParseRegexp", false, nil)
	}
	return c.validate("StringValue.MatchString", r.MatchString(string(c.val)), nil)
}

// MatchRegexp verify that the string matches the regular expression re
func (c *StringValue[T]) MatchRegexp(re *regexp.Regexp) *StringValue[T] {
	return c.validate("StringValue.MatchRegexp", re.MatchString(string(c.val)), nil)
}

// IPv4 verify that the string is formatted for IPv4.
func (c *StringValue[T]) IPv4() *StringValue[T] {
	return c.validate("StringValue.IPv4", isIPv4(string(c.val)), nil)
}

// IPv6 verify that the string is formatted for IPv6.
func (c *StringValue[T]) IPv6() *StringValue[T] {
	return c.validate("StringValue.IPv6", isIPv6(string(c.val)), nil)
}

// URL verify that the string is formatted for URL.
func (c *StringValue[T]) URL() *StringValue[T] {
	return c.validate("StringValue.URL", isURL(string(c.val)), nil)
}

// Email verify that the string is formatted for email.
func (c *StringValue[T]) Email() *StringValue[T] {
	return c.validate("StringValue.Email", isEmail(string(c.val)), nil)
}

// Alphabet Check if the string consists of letters
func (c *StringValue[T]) Alphabet() *StringValue[T] {
	return c.validate("StringValue.Alphabet", reAlphabet.MatchString(string(c.val)), nil)
}

// Numeric Check if the string consists of numbers
func (c *StringValue[T]) Numeric() *StringValue[T] {
	return c.validate("StringValue.Numeric", reNumeric.MatchString(string(c.val)), nil)
}

// AlphabetNumeric Check the string consists of letters and numbers.
func (c *StringValue[T]) AlphabetNumeric() *StringValue[T] {
	return c.validate("StringValue.AlphabetNumeric", reAlphabetNumeric.MatchString(string(c.val)), nil)
}

// Base64 verify that the string is formatted for base64.
func (c *StringValue[T]) Base64() *StringValue[T] {
	_, err := base64.StdEncoding.DecodeString(string(c.val))
	return c.validate("StringValue.Base64", err == nil, nil)
}

// Hex verify that the string is formatted for hex.
func (c *StringValue[T]) Hex() *StringValue[T] {
	_, err := hex.DecodeString(string(c.val))
	return c.validate("StringValue.Hex", err == nil, nil)
}

// Lowercase verify the string consists of lowercase letters.
func (c *StringValue[T]) Lowercase() *StringValue[T] {
	return c.validate("StringValue.Lowercase", string(c.val) == strings.ToLower(string(c.val)), nil)
}

// Uppercase verify the string consists of uppercase letters.
func (c *StringValue[T]) Uppercase() *StringValue[T] {
	return c.validate("StringValue.Uppercase", string(c.val) == strings.ToUpper(string(c.val)), nil)
}

// Customize customized data validation
// @layout error message
// @f check function
func (c *StringValue[T]) Customize(messageId string, f func(T) bool) *StringValue[T] {
	return c.validate(messageId, f(c.val), nil)
}
