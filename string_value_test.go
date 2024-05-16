package passport

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestStringValue_Required(t *testing.T) {
	assert.Error(t, NewString("name", "").Required().Err())
	assert.Error(t, NewString("name", "0").LenGt(1).Required().Err())
	assert.Nil(t, NewString("name", "aha").Required().Err())
}

func TestStringValue_Gt(t *testing.T) {
	assert.Error(t, NewString("age", "1").LenGt(2).Err())
	assert.Error(t, NewString("age", "").Required().LenGt(2).Err())
	assert.Nil(t, NewString("age", "2").LenGt(0).Err())
}

func TestStringValue_Eq(t *testing.T) {
	assert.Nil(t, NewString("age", "123").LenEq(3).Err())
	assert.Error(t, NewString("age", "25").LenEq(6).Err())
}

func TestStringValue_Gte(t *testing.T) {
	assert.Error(t, NewString("age", "1").LenGte(2).Err())
	assert.Error(t, NewString("age", "").Required().LenGte(1).Err())
	assert.Nil(t, NewString("age", "2").LenGte(1).Err())
	assert.Nil(t, NewString("age", "1").LenGte(1).Err())
}

func TestStringValue_Lt(t *testing.T) {
	assert.Error(t, NewString("age", "2").LenLt(1).Err())
	assert.Error(t, NewString("age", "").Required().LenLt(1).Err())
	assert.Nil(t, NewString("age", "1").LenLt(2).Err())
}

func TestStringValue_Lte(t *testing.T) {
	assert.Error(t, NewString("age", "2").LenLte(0).Err())
	assert.Error(t, NewString("age", "").Required().LenLte(1).Err())
	assert.Nil(t, NewString("age", "1").LenLte(2).Err())
	assert.Nil(t, NewString("age", "1").LenLte(1).Err())
}

func TestStringValue_Include(t *testing.T) {
	assert.Error(t, NewString("age", "2").IncludeBy("1", "3", "5").Err())
	assert.Error(t, NewString("age", "").Required().IncludeBy("1", "3", "5").Err())
	assert.Nil(t, NewString("age", "3").IncludeBy("1", "3", "5").Err())
}

func TestStringValue_Exclude(t *testing.T) {
	assert.Nil(t, NewString("age", "2").ExcludeBy("1", "3", "5").Err())
	assert.Error(t, NewString("age", "").Required().ExcludeBy("1", "3", "5").Err())
	assert.Error(t, NewString("age", "3").ExcludeBy("1", "3", "5").Err())
}

func TestStringValue_Email(t *testing.T) {
	assert.Nil(t, NewString("mail", "abc@qq.com").Email().Err())
	assert.Error(t, NewString("mail", "abc.qq").Email().Err())
	assert.Error(t, NewString("mail", "abc <abc@qq.com>").Email().Err())
}

func TestStringValue_IPv4(t *testing.T) {
	assert.Nil(t, NewString("ip", "192.168.1.1").IPv4().Err())
	assert.Error(t, NewString("ip", "256.168.1.1").IPv4().Err())
	assert.Error(t, NewString("ip", "192.168.1").IPv4().Err())
}

func TestStringValue_IPv6(t *testing.T) {
	assert.Nil(t, NewString("ip", "2001:0:2851:b9f0:2488:f0ba:210f:f3c8").IPv6().Err())
	assert.Error(t, NewString("ip", "2001:0:2851:b9f0:2488:f0ba:210f/f3c8").IPv6().Err())
}

func TestStringValue_URL(t *testing.T) {
	assert.Nil(t, NewString("url", "https://baidu.com").URL().Err())
	assert.Nil(t, NewString("url", "ws://baidu.com").URL().Err())
	assert.Error(t, NewString("url", "baidu.com").URL().Err())
	assert.Error(t, NewString("url", "https:///baidu.com/%x").URL().Err())
}

func TestStringValue_Alphabet(t *testing.T) {
	assert.Nil(t, NewString("name", "abc").Alphabet().Err())
	assert.Nil(t, NewString("name", "Na").Alphabet().Err())
	assert.Error(t, NewString("name", "Na2SO3").Alphabet().Err())
}

func TestStringValue_Numeric(t *testing.T) {
	assert.Error(t, NewString("name", "abc").Numeric().Err())
	assert.Nil(t, NewString("name", "123").Numeric().Err())
	assert.Error(t, NewString("name", "Na2SO3").Numeric().Err())
}

func TestStringValue_AlphabetNumeric(t *testing.T) {
	assert.Nil(t, NewString("name", "abc").AlphabetNumeric().Err())
	assert.Nil(t, NewString("name", "Na").AlphabetNumeric().Err())
	assert.Nil(t, NewString("name", "Na2SO3").AlphabetNumeric().Err())
	assert.Error(t, NewString("name", "Na/2SO3").AlphabetNumeric().Err())
}

func TestStringValue_Base64(t *testing.T) {
	assert.Error(t, NewString("name", "abc/@").Base64().Err())
	assert.Nil(t, NewString("name", base64.StdEncoding.EncodeToString([]byte{1, 2, 34})).Base64().Err())
}

func TestStringValue_Hex(t *testing.T) {
	assert.Nil(t, NewString("name", "0c1234").Hex().Err())
	assert.Nil(t, NewString("name", hex.EncodeToString([]byte{1, 2, 34})).Hex().Err())
	assert.Error(t, NewString("name", "0xyz").Hex().Err())
}

func TestStringValue_Lowercase(t *testing.T) {
	assert.Nil(t, NewString("name", "abc").Lowercase().Err())
	assert.Error(t, NewString("name", "Na").Lowercase().Err())
}

func TestStringValue_Uppercase(t *testing.T) {
	assert.Nil(t, NewString("name", "ABC").Uppercase().Err())
	assert.Error(t, NewString("name", "Na").Uppercase().Err())
}

func TestStringValue_MatchString(t *testing.T) {
	assert.Nil(t, NewString("name", "ABC").MatchString(`^[A-Z]+$`).Err())
	assert.Error(t, NewString("name", "ABCd").MatchString(`^[A-Z]+$`).Err())
	assert.Error(t, NewString("name", string([]byte{})).Required().MatchString(`^[A-Z]+$`).Err())

	t.Run("", func(t *testing.T) {
		var flateTail = []byte{0x00, 0x00, 0xff, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff}
		assert.Error(t, NewString("name", "abc").MatchString(string(flateTail)).Err())
	})
}

func TestStringValue_MatchRegexp(t *testing.T) {
	assert.Nil(t, NewString("name", "ABC").MatchRegexp(regexp.MustCompile(`^[A-Z]+$`)).Err())
	assert.Error(t, NewString("name", "ABCd").MatchRegexp(regexp.MustCompile(`^[A-Z]+$`)).Err())
	assert.Error(t, NewString("name", string([]byte{})).Required().MatchRegexp(regexp.MustCompile(`^[A-Z]+$`)).Err())
}
