package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExists(t *testing.T) {
	assert.True(t, PathExists("/etc/hosts"))
	assert.False(t, PathExists("/etc/hostsasdf"))
}
