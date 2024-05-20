package passport

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

var (
	reAlphabet        = regexp.MustCompile(`^[A-Za-z]+$`)
	reNumeric         = regexp.MustCompile(`^[0-9]+$`)
	reAlphabetNumeric = regexp.MustCompile(`^[A-Za-z0-9]+$`)
)

type StringValue[T ~string] struct {
	err error
	key string
	val T
}

func String[T ~string](k string, v T) *StringValue[T] {
	return &StringValue[T]{
		key: k,
		val: v,
	}
}

func (c *StringValue[T]) validate(ok bool, layout string, args ...any) *StringValue[T] {
	if c.err != nil || ok {
		return c
	}
	c.err = fmt.Errorf(layout, args...)
	return c
}

// Err get error
func (c *StringValue[T]) Err() error {
	return c.err
}

// Required the string cannot be empty
func (c *StringValue[T]) Required() *StringValue[T] {
	return c.validate(!isZero(c.val), "%s is required", c.key)
}

// Eq check that the string length is equal to v
func (c *StringValue[T]) Eq(v int) *StringValue[T] {
	return c.validate(len(c.val) == v, "Gth of %s should equal %v", c.key, v)
}

// Gt check that the string length is greater than v
func (c *StringValue[T]) Gt(v int) *StringValue[T] {
	return c.validate(len(c.val) > v, "Gth of %s should great than %v", c.key, v)
}

// Gte check that the string length is greater or equal than v
func (c *StringValue[T]) Gte(v int) *StringValue[T] {
	return c.validate(len(c.val) >= v, "Gth of %s should great or equal than %v", c.key, v)
}

// Lt check that the string length is less than v
func (c *StringValue[T]) Lt(v int) *StringValue[T] {
	return c.validate(len(c.val) < v, "Gth of %s should less than %v", c.key, v)
}

// Lte check that the string length is less or equal than v
func (c *StringValue[T]) Lte(v int) *StringValue[T] {
	return c.validate(len(c.val) <= v, "Gth of %s should less or equal than %v", c.key, v)
}

// IncludeBy check if args contains the string.
func (c *StringValue[T]) IncludeBy(args ...T) *StringValue[T] {
	return c.validate(contains(args, c.val), "%s should be one of %v", c.key, args)
}

// ExcludeBy checks if args does not contain the string.
func (c *StringValue[T]) ExcludeBy(args ...T) *StringValue[T] {
	return c.validate(!contains(args, c.val), "%s should not be one of %v", c.key, args)
}

// MatchString verify that the string matches the regular expression re
func (c *StringValue[T]) MatchString(re string) *StringValue[T] {
	if c.err != nil {
		return c
	}
	r, err := regexp.Compile(re)
	if err != nil {
		c.err = err
		return c
	}
	if !r.MatchString(string(c.val)) {
		c.err = fmt.Errorf("%s should match the regexp", c.key)
		return c
	}
	return c
}

// MatchRegexp verify that the string matches the regular expression re
func (c *StringValue[T]) MatchRegexp(re *regexp.Regexp) *StringValue[T] {
	return c.validate(re.MatchString(string(c.val)), "%s should match the regexp", c.key)
}

// IPv4 verify that the string is formatted for IPv4.
func (c *StringValue[T]) IPv4() *StringValue[T] {
	return c.validate(isIPv4(string(c.val)), "format of %s should be ipv4", c.key)
}

// IPv6 verify that the string is formatted for IPv6.
func (c *StringValue[T]) IPv6() *StringValue[T] {
	return c.validate(isIPv6(string(c.val)), "format of %s should be ipv6", c.key)
}

// URL verify that the string is formatted for URL.
func (c *StringValue[T]) URL() *StringValue[T] {
	return c.validate(isURL(string(c.val)), "format of %s should be url", c.key)
}

// Email verify that the string is formatted for email.
func (c *StringValue[T]) Email() *StringValue[T] {
	return c.validate(isEmail(string(c.val)), "format of %s should be email", c.key)
}

// Alphabet Check if the string consists of letters
func (c *StringValue[T]) Alphabet() *StringValue[T] {
	return c.validate(reAlphabet.MatchString(string(c.val)), "format of %s should be alphabet", c.key)
}

// Numeric Check if the string consists of numbers
func (c *StringValue[T]) Numeric() *StringValue[T] {
	return c.validate(reNumeric.MatchString(string(c.val)), "format of %s should be numeric", c.key)
}

// AlphabetNumeric Check the string consists of letters and numbers.
func (c *StringValue[T]) AlphabetNumeric() *StringValue[T] {
	return c.validate(reAlphabetNumeric.MatchString(string(c.val)), "format of %s should be alphabet numeric", c.key)
}

// Base64 verify that the string is formatted for base64.
func (c *StringValue[T]) Base64() *StringValue[T] {
	_, err := base64.StdEncoding.DecodeString(string(c.val))
	return c.validate(err == nil, "format of %s should be base64", c.key)
}

// Hex verify that the string is formatted for hex.
func (c *StringValue[T]) Hex() *StringValue[T] {
	_, err := hex.DecodeString(string(c.val))
	return c.validate(err == nil, "format of %s should be hex", c.key)
}

// Lowercase verify the string consists of lowercase letters.
func (c *StringValue[T]) Lowercase() *StringValue[T] {
	return c.validate(string(c.val) == strings.ToLower(string(c.val)), "format of %s should be lower case", c.key)
}

// Uppercase verify the string consists of uppercase letters.
func (c *StringValue[T]) Uppercase() *StringValue[T] {
	return c.validate(string(c.val) == strings.ToUpper(string(c.val)), "format of %s should be upper case", c.key)
}
