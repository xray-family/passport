package passport

import (
	"errors"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"net/http"
	"testing"
)

func TestPointerValue_Required(t *testing.T) {
	var r *http.Request
	assert.Error(t, Pointer("req", r).Required().Err())
	assert.Error(t, (&PointerValue[http.Request]{err: errors.New("test")}).Required().Err())
	assert.Nil(t, Pointer("req", new(http.Request)).Required().Err())

	t.Run("", func(t *testing.T) {
		var err = NewValidator(nil).Validate(
			Pointer[*http.Request]("名字", nil).Required(),
		)
		assert.Equal(t, err.Error(), "名字 cannot be empty")
	})

	t.Run("", func(t *testing.T) {
		var value = Pointer[*http.Request]("name", nil).Required()
		value.locConf.MessageID = "aha"
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		_ = GetBundle().AddMessages(Chinese, &i18n.Message{
			ID:    "Name",
			Other: "名字",
		})
		var r *http.Request
		err := NewValidator(nil).Validate(
			Pointer("Name", r).Required(),
		)
		assert.Equal(t, err.Error(), "Name cannot be empty")
	})
}

func TestPointerValue_Customize(t *testing.T) {
	tag := language.Make("en-US")
	message := &i18n.Message{
		ID:    "Customize",
		Other: "未成年人禁止入内",
	}
	_ = GetBundle().AddMessages(tag, message)

	t.Run("", func(t *testing.T) {
		var msg = Pointer("age", &http.Request{}).Customize(message.ID, func(i *http.Request) bool {
			return i == nil
		}).Err().Error()
		assert.Equal(t, msg, message.Other)
	})

	t.Run("", func(t *testing.T) {
		var err = Pointer("age", &http.Request{}).Customize(message.ID, func(i *http.Request) bool {
			return i != nil
		}).Err()
		assert.Nil(t, err)
	})
}
