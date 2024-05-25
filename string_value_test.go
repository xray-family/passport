package passport

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"regexp"
	"testing"
)

func TestStringValue_Required(t *testing.T) {
	assert.Error(t, String("name", "").Required().Err())
	assert.Error(t, String("name", "0").Gt(1).Required().Err())
	assert.Nil(t, String("name", "aha").Required().Err())
	assert.Error(t, String("name", " ").Required().Err())
	assert.Error(t, String("name", "\n").Required().Err())

	t.Run("", func(t *testing.T) {
		var err = NewValidator(WithLang("zh-CN")).Validate(
			String("名字", "").Required(),
		)
		assert.Equal(t, err.Error(), "名字不能为空")
	})

	t.Run("", func(t *testing.T) {
		_ = GetBundle().AddMessages(Chinese, &i18n.Message{
			ID:    "Name",
			Other: "名字",
		})
		err := NewValidator(WithAutoTranslate(true), WithLang(Chinese.String())).Validate(
			String("Name", "").Required(),
		)
		assert.Equal(t, err.Error(), "名字不能为空")
	})
}

func TestStringValue_Gt(t *testing.T) {
	assert.Error(t, String("age", "1").Gt(2).Err())
	assert.Error(t, String("age", "").Required().Gt(2).Err())
	assert.Nil(t, String("age", "2").Gt(0).Err())
}

func TestStringValue_Eq(t *testing.T) {
	assert.Nil(t, String("age", "123").Eq(3).Err())
	assert.Error(t, String("age", "25").Eq(6).Err())
}

func TestStringValue_Gte(t *testing.T) {
	assert.Error(t, String("age", "1").Gte(2).Err())
	assert.Error(t, String("age", "").Required().Gte(1).Err())
	assert.Nil(t, String("age", "2").Gte(1).Err())
	assert.Nil(t, String("age", "1").Gte(1).Err())
}

func TestStringValue_Lt(t *testing.T) {
	assert.Error(t, String("age", "2").Lt(1).Err())
	assert.Error(t, String("age", "").Required().Lt(1).Err())
	assert.Nil(t, String("age", "1").Lt(2).Err())
}

func TestStringValue_Lte(t *testing.T) {
	assert.Error(t, String("age", "2").Lte(0).Err())
	assert.Error(t, String("age", "").Required().Lte(1).Err())
	assert.Nil(t, String("age", "1").Lte(2).Err())
	assert.Nil(t, String("age", "1").Lte(1).Err())
}

func TestStringValue_Include(t *testing.T) {
	assert.Error(t, String("age", "2").In("1", "3", "5").Err())
	assert.Error(t, String("age", "").Required().In("1", "3", "5").Err())
	assert.Nil(t, String("age", "3").In("1", "3", "5").Err())
}

func TestStringValue_Email(t *testing.T) {
	assert.Nil(t, String("mail", "abc@qq.com").Email().Err())
	assert.Error(t, String("mail", "abc.qq").Email().Err())
	assert.Error(t, String("mail", "abc <abc@qq.com>").Email().Err())
}

func TestStringValue_IPv4(t *testing.T) {
	assert.Nil(t, String("ip", "192.168.1.1").IPv4().Err())
	assert.Error(t, String("ip", "256.168.1.1").IPv4().Err())
	assert.Error(t, String("ip", "192.168.1").IPv4().Err())
	assert.Error(t, String("ip", "2001:0:2851:b9f0:2488:f0ba:210f:f3c8").IPv4().Err())
}

func TestStringValue_IPv6(t *testing.T) {
	assert.Nil(t, String("ip", "2001:0:2851:b9f0:2488:f0ba:210f:f3c8").IPv6().Err())
	assert.Error(t, String("ip", "2001:0:2851:b9f0:2488:f0ba:210f/f3c8").IPv6().Err())
	assert.Error(t, String("ip", "192.168.1.1").IPv6().Err())
}

func TestStringValue_URL(t *testing.T) {
	assert.Nil(t, String("url", "https://baidu.com").URL().Err())
	assert.Nil(t, String("url", "ws://baidu.com").URL().Err())
	assert.Error(t, String("url", "baidu.com").URL().Err())
	assert.Error(t, String("url", "https:///baidu.com/%x").URL().Err())
}

func TestStringValue_Alphabet(t *testing.T) {
	assert.Nil(t, String("name", "abc").Alphabet().Err())
	assert.Nil(t, String("name", "Na").Alphabet().Err())
	assert.Error(t, String("name", "Na2SO3").Alphabet().Err())
}

func TestStringValue_Numeric(t *testing.T) {
	assert.Error(t, String("name", "abc").Numeric().Err())
	assert.Nil(t, String("name", "123").Numeric().Err())
	assert.Error(t, String("name", "Na2SO3").Numeric().Err())
}

func TestStringValue_AlphabetNumeric(t *testing.T) {
	assert.Nil(t, String("name", "abc").AlphabetNumeric().Err())
	assert.Nil(t, String("name", "Na").AlphabetNumeric().Err())
	assert.Nil(t, String("name", "Na2SO3").AlphabetNumeric().Err())
	assert.Error(t, String("name", "Na/2SO3").AlphabetNumeric().Err())
}

func TestStringValue_Base64(t *testing.T) {
	assert.Error(t, String("name", "abc/@").Base64().Err())
	assert.Nil(t, String("name", base64.StdEncoding.EncodeToString([]byte{1, 2, 34})).Base64().Err())
}

func TestStringValue_Hex(t *testing.T) {
	assert.Nil(t, String("name", "0c1234").Hex().Err())
	assert.Nil(t, String("name", hex.EncodeToString([]byte{1, 2, 34})).Hex().Err())
	assert.Error(t, String("name", "0xyz").Hex().Err())
}

func TestStringValue_Lowercase(t *testing.T) {
	assert.Nil(t, String("name", "abc").Lowercase().Err())
	assert.Error(t, String("name", "Na").Lowercase().Err())
}

func TestStringValue_Uppercase(t *testing.T) {
	assert.Nil(t, String("name", "ABC").Uppercase().Err())
	assert.Error(t, String("name", "Na").Uppercase().Err())
}

func TestStringValue_MatchString(t *testing.T) {
	assert.Nil(t, String("name", "ABC").MatchString(`^[A-Z]+$`).Err())
	assert.Error(t, String("name", "ABCd").MatchString(`^[A-Z]+$`).Err())
	assert.Error(t, String("name", string([]byte{})).Required().MatchString(`^[A-Z]+$`).Err())

	t.Run("", func(t *testing.T) {
		var flateTail = []byte{0x00, 0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff}
		assert.Error(t, String("name", "abc").MatchString(string(flateTail)).Err())
	})

	t.Run("", func(t *testing.T) {
		var value = String("name", "").Required()
		assert.Error(t, value.Err())
		assert.Error(t, value.Err())
	})
}

func TestStringValue_MatchRegexp(t *testing.T) {
	assert.Nil(t, String("name", "ABC").MatchRegexp(regexp.MustCompile(`^[A-Z]+$`)).Err())
	assert.Error(t, String("name", "ABCd").MatchRegexp(regexp.MustCompile(`^[A-Z]+$`)).Err())
	assert.Error(t, String("name", string([]byte{})).Required().MatchRegexp(regexp.MustCompile(`^[A-Z]+$`)).Err())
}

func TestStringValue_Customize(t *testing.T) {
	_ = GetBundle().AddMessages(language.Make("en-US"), &i18n.Message{
		ID:    "Customize",
		Other: "莫要喧哗",
	})
	t.Run("", func(t *testing.T) {
		const notice = "莫要喧哗"
		var msg = String("age", "1").Customize("Customize", func(s string) bool {
			return false
		}).Err().Error()
		assert.Equal(t, msg, notice)
	})

	t.Run("", func(t *testing.T) {
		const notice = "莫要喧哗"
		var err = String("age", "2").Customize("Customize", func(s string) bool {
			return true
		}).Err()
		assert.Nil(t, err)
	})
}

func TestStringValue_Between(t *testing.T) {
	assert.Nil(t, String("age", "3").Between("3", "5").Err())
	assert.Nil(t, String("age", "4").Between("3", "5").Err())
	assert.Error(t, String("age", "5").Between("3", "5").Err())
}
