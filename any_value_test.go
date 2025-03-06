package validator

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyValue_Customize(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var err = Any("name", 1).
			Customize("", func(a int) bool {
				return true
			}).
			Err()
		assert.Nil(t, err)
	})

	t.Run("", func(t *testing.T) {
		var err = Any("name", 1).
			Customize("", func(a int) bool {
				return false
			}).
			Err()
		assert.Error(t, err)
	})

	t.Run("", func(t *testing.T) {
		var err = Any("name", 1).
			Customize("", func(a int) bool {
				return false
			}).
			Customize("", func(a int) bool {
				return true
			}).
			Err()
		assert.Error(t, err)
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("", func(a int) bool {
				return false
			})
		value.Err()
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("", func(a int) bool {
				return false
			})
		//value.setConf(nil)
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("", func(a int) bool {
				return false
			})
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("StringValue.Required", func(a int) bool {
				return false
			})
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		_ = GetBundle().AddMessages(Chinese, &i18n.Message{
			ID:    "Name",
			Other: "名字",
		})
		err := NewValidator(newReq("zh-CN")).Validate(
			Any("Name", "").Customize("StringValue.Required", func(a string) bool {
				return false
			}),
		)
		assert.Equal(t, err.Error(), "Name 不能为空")
	})
}
