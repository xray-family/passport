package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceValue_Required(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, NewSlice("name", a1).Required().Err())
	assert.Error(t, NewSlice("name", a2).Required().Err())
	assert.Error(t, NewSlice("name", a2).LenGt(4).Required().Err())
}

func TestSliceValue_Eq(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, NewSlice("age", a1).LenEq(3).Err())
	assert.Error(t, NewSlice("age", a1).LenEq(4).Err())
	assert.Error(t, NewSlice("age", a2).Required().LenEq(0).Err())
}

func TestSliceValue_Gt(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, NewSlice("age", a1).LenGt(2).Err())
	assert.Error(t, NewSlice("age", a1).LenGt(4).Err())
	assert.Error(t, NewSlice("age", a2).Required().LenGt(0).Err())
}

func TestSliceValue_Gte(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, NewSlice("age", a1).LenGte(2).Err())
	assert.Nil(t, NewSlice("age", a1).LenGte(3).Err())
	assert.Error(t, NewSlice("age", a1).LenGte(4).Err())
	assert.Error(t, NewSlice("age", a2).Required().LenGte(0).Err())
}

func TestSliceValue_Lt(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, NewSlice("age", a1).LenLt(4).Err())
	assert.Error(t, NewSlice("age", a1).LenLt(1).Err())
	assert.Error(t, NewSlice("age", a2).Required().LenLt(4).Err())
}

func TestSliceValue_Lte(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, NewSlice("age", a1).LenLte(4).Err())
	assert.Nil(t, NewSlice("age", a1).LenLte(3).Err())
	assert.Error(t, NewSlice("age", a1).LenLte(1).Err())
	assert.Error(t, NewSlice("age", a2).Required().LenLte(4).Err())
}

func TestSliceValue_Include(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Error(t, NewSlice("age", a1).Include(2).Err())
	assert.Error(t, NewSlice("age", a2).Required().Include(1).Err())
	assert.Nil(t, NewSlice("age", a1).Include(1).Err())
}

func TestSliceValue_Exclude(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Error(t, NewSlice("age", a1).Exclude(1).Err())
	assert.Error(t, NewSlice("age", a2).Required().Exclude(1).Err())
	assert.Nil(t, NewSlice("age", a1).Exclude(2).Err())
}
