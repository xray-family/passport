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
}
