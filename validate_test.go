package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	GetBundle()
	GetLocalizer()

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
		err := NewValidator(WithLang("zh-CN")).Validate(
			Ordered("Name", r.Name).Required(),
			Ordered("Age", r.Age).Gte(18),
		)
		assert.Equal(t, err.Error(), "Age必须大于等于18")
	})
}
