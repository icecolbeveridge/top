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
	// t.Log(grid)
	x := contour.Contours(grid, -5, -4, -3, -2, -1, -0.5, 0, 0.5, 1, 2, 3, 4, 5)
	ff, err := os.Create("/tmp/contour_test.png")
	assert.NoError(err)
	contour.Plot(x, ff)
}
