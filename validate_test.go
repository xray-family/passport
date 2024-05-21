package passport

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestValidate(t *testing.T) {
	SetLang("en-US")

	type Req struct {
		Name string
		Age  int
	}

	t.Run("", func(t *testing.T) {
		r := Req{Name: "aha", Age: 3}
		err := Validate(
			Ordered("Name", r.Name).Required(),
			Ordered("Age", r.Age).Gte(18),
		)
		assert.Error(t, err)
	})

	t.Run("", func(t *testing.T) {
		r := Req{Name: "aha", Age: 20}
		err := Validate(
			Ordered("Name", r.Name).Required(),
			Ordered("Age", r.Age).Gte(18),
		)
		assert.Nil(t, err)
	})

	t.Run("", func(t *testing.T) {
		r := Req{Name: "aha", Age: 20}
		err := NewValidator().Validate(
			Ordered("Name", r.Name).Required(),
			Ordered("Age", r.Age).Gte(18),
		)
		assert.Nil(t, err)
	})

	t.Run("", func(t *testing.T) {
		r := Req{Name: "aha", Age: 15}
		h := http.Header{}
		h.Set("Accept-Language", "zh-CN")
		err := NewValidatorFromHTTP(h).Validate(
			Ordered("Name", r.Name).Required(),
			Ordered("Age", r.Age).Gte(18),
		)
		assert.Equal(t, err.Error(), "Age必须大于等于18")
	})
}
