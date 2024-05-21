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
	err        error
	key        string
	val        T
	mark       bool
	localizer  *i18n.Localizer
	config     *i18n.LocalizeConfig
	messageMap map[string]*i18n.Message
}

func String[T ~string](k string, v T) *StringValue[T] {
	return &StringValue[T]{
		key: k,
		val: v,
	}
}

func (c *StringValue[T]) setLocalizer(localizer *i18n.Localizer) {
	c.localizer = localizer
}

func (c *StringValue[T]) validate(messageId string, ok bool, val any) *StringValue[T] {
	if c.mark || ok {
		return c
	}
	c.mark = true
	messageId = "StringValue." + messageId
	c.config = &i18n.LocalizeConfig{
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

	if c.localizer == nil {
		c.localizer = i18n.NewLocalizer(_bundle, defaultLanguages...)
	}

	if len(c.messageMap) > 0 {
		if message := c.messageMap[c.config.MessageID]; message != nil {
			c.config.DefaultMessage = message
			c.config.MessageID, c.config.DefaultMessage.ID = "!undefined", "!undefined"
			str, _ := c.localizer.Localize(c.config)
			c.err = errors.New(str)
			return c.err
		}
	}

	str, err := c.localizer.Localize(c.config)
	if err != nil {
		c.err = err
		return c.err
	}
	c.err = errors.New(str)
	return c.err
}

// Required the string cannot be empty
func (c *StringValue[T]) Required() *StringValue[T] {
	return c.validate("Required", !isZero(c.val), nil)
}

// Eq check that the string length is equal to v
func (c *StringValue[T]) Eq(v int) *StringValue[T] {
	return c.validate("Eq", len(c.val) == v, v)
}

// Gt check that the string length is greater than v
func (c *StringValue[T]) Gt(v int) *StringValue[T] {
	return c.validate("Gt", len(c.val) > v, v)
}

// Gte check that the string length is greater or equal than v
func (c *StringValue[T]) Gte(v int) *StringValue[T] {
	return c.validate("gte", len(c.val) >= v, v)
}

// Lt check that the string length is less than v
func (c *StringValue[T]) Lt(v int) *StringValue[T] {
	return c.validate("Lt", len(c.val) < v, v)
}

// Lte check that the string length is less or equal than v
func (c *StringValue[T]) Lte(v int) *StringValue[T] {
	return c.validate("Lte", len(c.val) <= v, v)
}

// IncludeBy check if args contains the string.
func (c *StringValue[T]) IncludeBy(args ...T) *StringValue[T] {
	return c.validate("IncludeBy", contains(args, c.val), args)
}

// ExcludeBy checks if args does not contain the string.
func (c *StringValue[T]) ExcludeBy(args ...T) *StringValue[T] {
	return c.validate("ExcludeBy", !contains(args, c.val), args)
}

// MatchString verify that the string matches the regular expression re
func (c *StringValue[T]) MatchString(re string) *StringValue[T] {
	if c.mark {
		return c
	}
	r, err := regexp.Compile(re)
	if err != nil {
		return c.validate("ParseRegexp", false, nil)
	}
	return c.validate("MatchString", r.MatchString(string(c.val)), nil)
}

// MatchRegexp verify that the string matches the regular expression re
func (c *StringValue[T]) MatchRegexp(re *regexp.Regexp) *StringValue[T] {
	return c.validate("MatchRegexp", re.MatchString(string(c.val)), nil)
}

// IPv4 verify that the string is formatted for IPv4.
func (c *StringValue[T]) IPv4() *StringValue[T] {
	return c.validate("IPv4", isIPv4(string(c.val)), nil)
}

// IPv6 verify that the string is formatted for IPv6.
func (c *StringValue[T]) IPv6() *StringValue[T] {
	return c.validate("IPv6", isIPv6(string(c.val)), nil)
}

// URL verify that the string is formatted for URL.
func (c *StringValue[T]) URL() *StringValue[T] {
	return c.validate("URL", isURL(string(c.val)), nil)
}

// Email verify that the string is formatted for email.
func (c *StringValue[T]) Email() *StringValue[T] {
	return c.validate("Email", isEmail(string(c.val)), nil)
}

// Alphabet Check if the string consists of letters
func (c *StringValue[T]) Alphabet() *StringValue[T] {
	return c.validate("Alphabet", reAlphabet.MatchString(string(c.val)), nil)
}

// Numeric Check if the string consists of numbers
func (c *StringValue[T]) Numeric() *StringValue[T] {
	return c.validate("Numeric", reNumeric.MatchString(string(c.val)), nil)
}

// AlphabetNumeric Check the string consists of letters and numbers.
func (c *StringValue[T]) AlphabetNumeric() *StringValue[T] {
	return c.validate("AlphabetNumeric", reAlphabetNumeric.MatchString(string(c.val)), nil)
}

// Base64 verify that the string is formatted for base64.
func (c *StringValue[T]) Base64() *StringValue[T] {
	_, err := base64.StdEncoding.DecodeString(string(c.val))
	return c.validate("Base64", err == nil, nil)
}

// Hex verify that the string is formatted for hex.
func (c *StringValue[T]) Hex() *StringValue[T] {
	_, err := hex.DecodeString(string(c.val))
	return c.validate("Hex", err == nil, nil)
}

// Lowercase verify the string consists of lowercase letters.
func (c *StringValue[T]) Lowercase() *StringValue[T] {
	return c.validate("Lowercase", string(c.val) == strings.ToLower(string(c.val)), nil)
}

// Uppercase verify the string consists of uppercase letters.
func (c *StringValue[T]) Uppercase() *StringValue[T] {
	return c.validate("Uppercase", string(c.val) == strings.ToUpper(string(c.val)), nil)
}

// Customize customized data validation
// @layout error message
// @f check function
func (c *StringValue[T]) Customize(layout string, f func(T) bool) *StringValue[T] {
	const funcName = "Customize"
	return c.validate(funcName, f(c.val), nil).Message(funcName, layout)
}

// Message customizing error messages
func (c *StringValue[T]) Message(funcName string, layout string) *StringValue[T] {
	if c.messageMap == nil {
		c.messageMap = make(map[string]*i18n.Message)
	}
	id := "StringValue." + funcName
	c.messageMap[id] = &i18n.Message{
		ID:    id,
		Other: layout,
	}
	return c
}
