package contour_test

import (
	"testing"
	"os"
	top_io "top/io"
	"top/contour"
	"github.com/stretchr/testify/assert"
)

const ROOT = "/home/colin/go/top" // TODO: global util


func TestContour(t *testing.T) {
	assert := assert.New(t)
	f, err := os.Open(ROOT + "/data/example.pns")
	assert.NoError(err)

	T := top_io.ReadPNS(f)
	t.Log(T)

	t.Log(contour.Autogrid(T))
}
