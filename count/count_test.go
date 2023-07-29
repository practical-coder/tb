package count

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromReader(t *testing.T) {
	str := "test string\nlala\nlala\nlalalalalalalalalalalalalalalala"
	r := strings.NewReader(str)
	c, err := FromReader(r)
	if assert.Nil(t, err) {
		assert.Equal(t, len(str), c)
	}
}
