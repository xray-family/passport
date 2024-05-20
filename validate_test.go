package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
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
}

func TestValidateErrors(t *testing.T) {
	type Req struct {
		Name string
		Age  int
	}

	t.Run("", func(t *testing.T) {
		r := Req{Name: "aha", Age: 3}
		err := ValidateErrors(
			Ordered("Name", r.Name).Required().Err(),
			Ordered("Age", r.Age).Gte(18).Err(),
		)
		assert.Error(t, err)
	})

	t.Run("", func(t *testing.T) {
		r := Req{Name: "aha", Age: 20}
		err := ValidateErrors(
			Ordered("Name", r.Name).Required().Err(),
			Ordered("Age", r.Age).Gte(18).Err(),
		)
		assert.Nil(t, err)
	})
}
