package golib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExists(t *testing.T) {
	assert.True(t, Exists("/etc/hosts"))
	assert.False(t, Exists("/etc/hostsasdf"))
}
