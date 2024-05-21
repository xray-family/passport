package passport

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPointerValue_Required(t *testing.T) {
	var r *http.Request
	assert.Error(t, Pointer("req", r).Required().Err())
	assert.Error(t, (&PointerValue[http.Request]{err: errors.New("test")}).Required().Err())
	assert.Nil(t, Pointer("req", new(http.Request)).Required().Err())

	t.Run("", func(t *testing.T) {
		var err = NewValidator("zh-CN").Validate(
			Pointer[*http.Request]("名字", nil).Required(),
		)
		assert.Equal(t, err.Error(), "名字不能为空")
	})

	t.Run("", func(t *testing.T) {
		var value = Pointer[*http.Request]("name", nil).Required()
		value.config.MessageID = "aha"
		assert.Error(t, value.Err())
	})
}

func TestPointerValue_Customize(t *testing.T) {
	t.Run("", func(t *testing.T) {
		const notice = "不要大声喧哗"
		var msg = Pointer("age", &http.Request{}).Customize(notice, func(i *http.Request) bool {
			return i == nil
		}).Err().Error()
		assert.Equal(t, msg, notice)
	})

	t.Run("", func(t *testing.T) {
		const notice = "不要大声喧哗"
		var err = Pointer("age", &http.Request{}).Customize(notice, func(i *http.Request) bool {
			return i != nil
		}).Err()
		assert.Nil(t, err)
	})
}
