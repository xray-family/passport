package passport

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestSliceValue_Required(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, Slice("name", a1).Required().Err())
	assert.Error(t, Slice("name", a2).Required().Err())
	assert.Error(t, Slice("name", a2).Gt(4).Required().Err())

	t.Run("", func(t *testing.T) {
		var err = NewValidator(WithLang("zh-CN")).Validate(
			Slice("name", a2).Required(),
		)
		assert.Equal(t, err.Error(), "name不能为空")
	})

	t.Run("", func(t *testing.T) {
		var value = Slice("name", a2).Required()
		value.Err()
		assert.Equal(t, value.Err().Error(), "name cannot be empty")
	})

	t.Run("", func(t *testing.T) {
		var value = Slice("name", a2).Required()
		value.locConf.MessageID = "aha"
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		_ = GetBundle().AddMessages(Chinese, &i18n.Message{
			ID:    "Name",
			Other: "名字",
		})
		err := NewValidator(WithAutoTranslate(true), WithLang(Chinese.String())).Validate(
			Slice("Name", a2).Required(),
		)
		assert.Equal(t, err.Error(), "名字不能为空")
	})
}

func TestSliceValue_Eq(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, Slice("age", a1).Eq(3).Err())
	assert.Error(t, Slice("age", a1).Eq(4).Err())
	assert.Error(t, Slice("age", a2).Required().Eq(0).Err())
}

func TestSliceValue_Gt(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, Slice("age", a1).Gt(2).Err())
	assert.Error(t, Slice("age", a1).Gt(4).Err())
	assert.Error(t, Slice("age", a2).Required().Gt(0).Err())
}

func TestSliceValue_Gte(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, Slice("age", a1).Gte(2).Err())
	assert.Nil(t, Slice("age", a1).Gte(3).Err())
	assert.Error(t, Slice("age", a1).Gte(4).Err())
	assert.Error(t, Slice("age", a2).Required().Gte(0).Err())
}

func TestSliceValue_Lt(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, Slice("age", a1).Lt(4).Err())
	assert.Error(t, Slice("age", a1).Lt(1).Err())
	assert.Error(t, Slice("age", a2).Required().Lt(4).Err())
}

func TestSliceValue_Lte(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, Slice("age", a1).Lte(4).Err())
	assert.Nil(t, Slice("age", a1).Lte(3).Err())
	assert.Error(t, Slice("age", a1).Lte(1).Err())
	assert.Error(t, Slice("age", a2).Required().Lte(4).Err())
}

func TestSliceValue_Include(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Error(t, Slice("age", a1).Contains(2).Err())
	assert.Error(t, Slice("age", a2).Required().Contains(1).Err())
	assert.Nil(t, Slice("age", a1).Contains(1).Err())
}

func TestSliceValue_Customize(t *testing.T) {
	tag := language.Make("en-US")
	message := &i18n.Message{
		ID:    "Customize",
		Other: "未成年人禁止入内",
	}
	_ = GetBundle().AddMessages(tag, message)

	t.Run("", func(t *testing.T) {
		var arr []int
		var msg = Slice("age", arr).Customize(message.ID, func(i []int) bool {
			return len(i) >= 18
		}).Err().Error()
		assert.Equal(t, msg, message.Other)
	})

	t.Run("", func(t *testing.T) {
		var arr []int
		var err = Slice("age", arr).Customize(message.ID, func(i []int) bool {
			return len(i) >= 0
		}).Err()
		assert.Nil(t, err)
	})
}
