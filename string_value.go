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

func NewString[T ~string](k string, v T) *StringValue[T] {
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

func (c *StringValue[T]) Err() error {
	return c.err
}

func (c *StringValue[T]) Required() *StringValue[T] {
	return c.validate(!isZero(c.val), "%s is required", c.key)
}

func (c *StringValue[T]) LenEq(v int) *StringValue[T] {
	return c.validate(len(c.val) == v, "length of %s should equal %v", c.key, v)
}

func (c *StringValue[T]) LenGt(v int) *StringValue[T] {
	return c.validate(len(c.val) > v, "length of %s should great than %v", c.key, v)
}

func (c *StringValue[T]) LenGte(v int) *StringValue[T] {
	return c.validate(len(c.val) >= v, "length of %s should great or equal than %v", c.key, v)
}

func (c *StringValue[T]) LenLt(v int) *StringValue[T] {
	return c.validate(len(c.val) < v, "length of %s should less than %v", c.key, v)
}

func (c *StringValue[T]) LenLte(v int) *StringValue[T] {
	return c.validate(len(c.val) <= v, "length of %s should less or equal than %v", c.key, v)
}

func (c *StringValue[T]) IncludeBy(args ...T) *StringValue[T] {
	return c.validate(contains(args, c.val), "%s should be one of %v", c.key, args)
}

func (c *StringValue[T]) ExcludeBy(args ...T) *StringValue[T] {
	return c.validate(!contains(args, c.val), "%s should not be one of %v", c.key, args)
}

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

func (c *StringValue[T]) MatchRegexp(re *regexp.Regexp) *StringValue[T] {
	return c.validate(re.MatchString(string(c.val)), "%s should match the regexp", c.key)
}

func (c *StringValue[T]) IPv4() *StringValue[T] {
	return c.validate(isIPv4(string(c.val)), "format of %s should be ipv4", c.key)
}

func (c *StringValue[T]) IPv6() *StringValue[T] {
	return c.validate(isIPv6(string(c.val)), "format of %s should be ipv6", c.key)
}

func (c *StringValue[T]) URL() *StringValue[T] {
	return c.validate(isURL(string(c.val)), "format of %s should be url", c.key)
}

func (c *StringValue[T]) Email() *StringValue[T] {
	return c.validate(isEmail(string(c.val)), "format of %s should be email", c.key)
}

func (c *StringValue[T]) Alphabet() *StringValue[T] {
	return c.validate(reAlphabet.MatchString(string(c.val)), "format of %s should be alphabet", c.key)
}

func (c *StringValue[T]) Numeric() *StringValue[T] {
	return c.validate(reNumeric.MatchString(string(c.val)), "format of %s should be numeric", c.key)
}

func (c *StringValue[T]) AlphabetNumeric() *StringValue[T] {
	return c.validate(reAlphabetNumeric.MatchString(string(c.val)), "format of %s should be alphabet numeric", c.key)
}

func (c *StringValue[T]) Base64() *StringValue[T] {
	_, err := base64.StdEncoding.DecodeString(string(c.val))
	return c.validate(err == nil, "format of %s should be base64", c.key)
}

func (c *StringValue[T]) Hex() *StringValue[T] {
	_, err := hex.DecodeString(string(c.val))
	return c.validate(err == nil, "format of %s should be hex", c.key)
}

func (c *StringValue[T]) Lowercase() *StringValue[T] {
	return c.validate(string(c.val) == strings.ToLower(string(c.val)), "format of %s should be lower case", c.key)
}

func (c *StringValue[T]) Uppercase() *StringValue[T] {
	return c.validate(string(c.val) == strings.ToUpper(string(c.val)), "format of %s should be upper case", c.key)
}
