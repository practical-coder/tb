package read

import (
	"testing"

	"github.com/practical-coder/tb/count"
	"github.com/stretchr/testify/assert"
)

func TestFromWeb(t *testing.T) {
	rc, err := FromWeb("http://example.org/")
	assert.Nil(t, err)
	count, err := count.FromReader(rc)
	assert.Nil(t, err)
	assert.Equal(t, 1256, count)
}

func TestFromFile(t *testing.T) {
	rc, err := FromFile("abc.txt")
	assert.Nil(t, err)
	count, err := count.FromReader(rc)
	assert.Nil(t, err)
	assert.Equal(t, 4, count)
}
