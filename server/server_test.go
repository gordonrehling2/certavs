package server

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestServerWhereItExitsCleanly(t *testing.T) {
	exitCode := Server()
	assert.Equal(t, 0, exitCode)
}
