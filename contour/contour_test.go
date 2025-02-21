package contour_test

import (
	"os"
	"testing"
	"top/contour"
	top_io "top/io"

	"github.com/stretchr/testify/assert"
)

const ROOT = "/home/colin/go/top" // TODO: global util

func TestContour(t *testing.T) {
	assert := assert.New(t)
	f, err := os.Open(ROOT + "/data/example.pns")
	assert.NoError(err)

	T := top_io.ReadPNS(f)
	t.Log(T)
	grid := contour.Autogrid(T)
	x := contour.Contour(grid, 0)
	assert.NotEmpty(x)
}
