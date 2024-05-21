package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyValue_Customize(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var err = Any("name", 1).
			Customize("", func(a any) bool {
				return true
			}).
			Err()
		assert.Nil(t, err)
	})

	t.Run("", func(t *testing.T) {
		var err = Any("name", 1).
			Customize("", func(a any) bool {
				return false
			}).
			Err()
		assert.Error(t, err)
	})

	t.Run("", func(t *testing.T) {
		var err = Any("name", 1).
			Customize("", func(a any) bool {
				return false
			}).
			Customize("", func(a any) bool {
				return true
			}).
			Err()
		assert.Error(t, err)
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("", func(a any) bool {
				return false
			})
		value.Err()
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("", func(a any) bool {
				return false
			})
		value.setLang("en-US")
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("", func(a any) bool {
				return false
			})
		value.messageMap = nil
		assert.Error(t, value.Err())
	})

	t.Run("", func(t *testing.T) {
		var value = Any("name", 1).
			Customize("", func(a any) bool {
				return false
			})
		value.messageMap = nil
		value.config.MessageID = "aha"
		assert.Error(t, value.Err())
	})
}
