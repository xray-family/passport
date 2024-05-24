package passport

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestOrderedValue_Required(t *testing.T) {
	assert.Error(t, Ordered("name", "").Required().Err())
	assert.Error(t, Ordered("name", 0).Required().Err())
	assert.Error(t, Ordered("name", 0).Gt(1).Required().Err())
	assert.Nil(t, Ordered("name", "aha").Required().Err())

	t.Run("", func(t *testing.T) {
		_ = GetBundle().AddMessages(Chinese, &i18n.Message{
			ID:    "Name",
			Other: "名字",
		})
		err := NewValidator(WithAutoTranslate(), WithLang(Chinese.String())).Validate(
			Ordered("Name", "").Required(),
		)
		assert.Equal(t, err.Error(), "名字不能为空")
	})
}

func TestOrderedValue_Gt(t *testing.T) {
	assert.Error(t, Ordered("age", 1).Gt(2).Err())
	assert.Error(t, Ordered("age", 0).Required().Gt(2).Err())
	assert.Nil(t, Ordered("age", 2).Gt(1).Err())
}

func TestOrderedValue_Gte(t *testing.T) {
	assert.Error(t, Ordered("age", 1).Gte(2).Err())
	assert.Error(t, Ordered("age", 0).Required().Gte(1).Err())
	assert.Nil(t, Ordered("age", 2).Gte(1).Err())
	assert.Nil(t, Ordered("age", 1).Gte(1).Err())
}

func TestOrderedValue_Lt(t *testing.T) {
	assert.Error(t, Ordered("age", 2).Lt(1).Err())
	assert.Error(t, Ordered("age", 0).Required().Lt(1).Err())
	assert.Nil(t, Ordered("age", 1).Lt(2).Err())
}

func TestOrderedValue_Lte(t *testing.T) {
	assert.Error(t, Ordered("age", 2).Lte(1).Err())
	assert.Error(t, Ordered("age", 0).Required().Lte(1).Err())
	assert.Nil(t, Ordered("age", 1).Lte(2).Err())
	assert.Nil(t, Ordered("age", 1).Lte(1).Err())
}

func TestOrderedValue_Include(t *testing.T) {
	assert.Error(t, Ordered("age", 2).In(1, 3, 5).Err())
	assert.Error(t, Ordered("age", 0).Required().In(1, 3, 5).Err())
	assert.Nil(t, Ordered("age", 3).In(1, 3, 5).Err())
}

func TestOrderedValue_Err(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var validator = NewValidator(WithLang("en-US"))
		var v = Ordered("age", 1).Gt(18)
		v.setConf(validator.conf)
		assert.Error(t, v.Err())
		assert.Error(t, v.Err())
	})

	t.Run("", func(t *testing.T) {
		var v = Ordered("age", 1).Gt(18)
		v.locConf.MessageID = "oh"
		assert.Error(t, v.Err())
	})
}

func TestOrderedValue_Customize(t *testing.T) {
	tag := language.Make("en-US")
	message := &i18n.Message{
		ID:    "Customize",
		Other: "未成年人禁止入内",
	}
	_ = GetBundle().AddMessages(tag, message)
	t.Run("", func(t *testing.T) {
		var msg = Ordered("age", 1).Customize(message.ID, func(i int) bool {
			return i >= 18
		}).Err().Error()
		assert.Equal(t, msg, message.Other)
	})

	t.Run("", func(t *testing.T) {
		var err = Ordered("age", 20).Customize(message.ID, func(i int) bool {
			return i >= 18
		}).Err()
		assert.Nil(t, err)
	})
}

func TestOrderedValue_Between(t *testing.T) {
	assert.Nil(t, Ordered("age", 3).Between(3, 5).Err())
	assert.Nil(t, Ordered("age", 4).Between(3, 5).Err())
	assert.Error(t, Ordered("age", 5).Between(3, 5).Err())
}
