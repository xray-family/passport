package passport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceValue_Required(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Nil(t, Slice("name", a1).Required().Err())
	assert.Error(t, Slice("name", a2).Required().Err())
	assert.Error(t, Slice("name", a2).Gt(4).Required().Err())
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
	assert.Error(t, Slice("age", a1).Include(2).Err())
	assert.Error(t, Slice("age", a2).Required().Include(1).Err())
	assert.Nil(t, Slice("age", a1).Include(1).Err())
}

func TestSliceValue_Exclude(t *testing.T) {
	var a1 = []int{1, 3, 5}
	var a2 []int
	assert.Error(t, Slice("age", a1).Exclude(1).Err())
	assert.Error(t, Slice("age", a2).Required().Exclude(1).Err())
	assert.Nil(t, Slice("age", a1).Exclude(2).Err())
}
