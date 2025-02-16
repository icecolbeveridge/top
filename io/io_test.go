package top_io_test

import (
	"os"
	"testing"
	top_io "top/io"

	"github.com/stretchr/testify/assert"
)

const ROOT = "/home/colin/go/top" // TODO: global util

func TestReadPNS(t *testing.T) {
	assert := assert.New(t)

	f, err := os.Open(ROOT + "/data/ar930605.pns")
	assert.NoError(err)

	T := top_io.ReadPNS(f)
	assert.NotNil(T)
	t.Log(T)
}
