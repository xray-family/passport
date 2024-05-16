package passport

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPointerValue_Required(t *testing.T) {
	var r *http.Request
	assert.Error(t, NewPointer("req", r).Required().Err())
	assert.Error(t, (&PointerValue[http.Request]{err: errors.New("test")}).Required().Err())
	assert.Nil(t, NewPointer("req", new(http.Request)).Required().Err())
}
