package contour

import "top/top"

// This is a thinking-it-through piece.
//
// To develop a 3D contour surface, I think the plan would be:
//  * start with a grid in the plane
//  * find broken edges in a grid with alternating slants
//  * determine broken faces per tetrahedron
//  * add neighbouring tetrahedra to stack
//  * once broken tets are found, approximate the triangles
// 	 * (it's possible to have a quad, in which case, we can busk it)
//
// I'll need a method for converting from grid coords to pns-space and back again
//
// It's going to be messy, but it's going to be fun!
//
// Gridwise, I think the vertices of the cube are at (0,0,0) to (12,12,12)
// in the obvious sense.
// The edges of length 1 are at (6,[0|12], [0|12]) and permutations thereof.
// The sqrt(2) edges are at (6,6,[0|12]) and perms thereof.
// I'll need to think about the directions.
// The faces are... where are they? The orthogonal ones are centred at
// places like (0,4,4) -- the edges for that are (0,6,0), (0,0,6) and (0,6,6),
// which gives us an average.
// What about the slanted faces? Sample points are (0,0,0), (12,12,0), and (0,12,12), giving (4,8,4).
// And the tets? The one with points (0,0,0), (12,0,0), (12,12,0) and (12,0,12) has centroid (9, 3, 3).
// The big central tet is at (6,6,6).

// And then... what can we say about adjacent cubes?
// The outer tet centroids would be at:
//  - (9, 3, 3) , (3, 9, 3) ,  (3, 3, 9), (9, 9, 9) (which is kind of nice)
// Reflecting in x=12 gives:
//  - (15, 3, 3) , (21, 9, 3), (21, 3, 9), (15, 9, 9)
// [or (-9, 3, 3), (-3, 3, 3), (-3, 3, 9), (-9, 9, 9)]
// ... and we're going to need to do some book-keeping.
//
// Or are we? What are we trying to accomplish? We start with a grid
// of true/false in the x-y plane. We then use our criss-cross grid to
// identify split bases, which we put on the investigate list.
// We determine the midpoint facet for one of these, and add any adjacent tets to the
// "investigate" list (unless it's already handled).
// Then we pop from the list and repeat.
//
// Each broken diagonal edge abuts six tets --
// four differ by (b, a, -b) [assuming a is the cube face coord]
// and two by (0, 2a, 0) [similar]
// -- does this hold for all?
//
// Each orthog edge abuts four tets --
//  offset by +/-1 in cube edge directions, and away from "point" of tet
//
// I suspect I'm overthinking this. We can proceed a cube at a time.
//
// I'm also going to decouple "finding interesting cubes" from "finding the
// facets".
//
// So, the plan is:
//  -- seed Interesting with cubes having broken bases
//  -- take an Interesting but Unhandled cube:
//  	-- check for broken faces
// 		-- add any Interesting neighbours to Interesting
//  -- continue until every Interesting cube is Handled.
//
//  Looking at the .obj file format, this is great -- we can number the
// vertices and then figure out where they are.

type Fielder func(x, y, z float64) float64

type facet [3]int
type gridpoint [3]int
type Shell struct {
	Points []top.Point
	Facets []facet
}

type Contour3DOptions struct {
	xmin, xmax, ymin, ymax, zmin, zmax float64
	nx, ny, nz                         int
	fn                                 func(top.Vector) float64
	level                              float64 // probably a slice in the end TODO
}

func (c Contour3DOptions) gridToXYZ(gp gridpoint) top.Vector {
	x := c.xmin + float64(gp[0])*(c.xmax-c.xmin)/float64(c.nx-1)
	y := c.ymin + float64(gp[1])*(c.ymax-c.ymin)/float64(c.ny-1)
	z := c.zmin + float64(gp[2])*(c.zmax-c.zmin)/float64(c.nz-1)
	return top.Vector{X: x, Y: y, Z: z}
}

type ccube struct {
	gp    gridpoint
	field [2][2][2]float64
}

// TODO: options
func Contour3d(c Contour3DOptions) []Shell {
	out := make([]Shell, 0)
	grid := make(map[gridpoint]float64)
	// start by filling the z=0 grid
	for x := 0; x < c.nx; x++ {
		for y := 0; y < c.ny; y++ {
			gp := gridpoint{x, y, 0}
			v := c.gridToXYZ(gp)
			grid[gp] = c.fn(v)
		}
	}
	// any grid cell with a mixture of above/below is interesting
	to_process := make([]gridpoint, 0) // index cubes by minimal coords
	interesting := make(map[gridpoint]ccube)
	uninteresting := make(map[gridpoint]ccube)

	// seed interesting with z=0 interesting cubes
	for x := 0; x < c.nx; x++ {
		for y := 0; y < c.ny; y++ {
			cell_total := 0
			for i := 0; i <= 1; i++ {
				for j := 0; j <= 1; j++ {
					gp := gridpoint{x + i, y + j, 0}
					if grid[gp] > c.level {
						cell_total += 1
					}
				}
			}
			if cell_total > 0 && cell_total < 4 {
				to_process = append(to_process, gridpoint{x, y, 0})
			}
		}
	}
	// process each interesting cell and add any interesting neighbours TODO
	var current gridpoint
	for len(to_process) > 0 {
		l := len(to_process)
		to_process, current = to_process[:l-1], to_process[l-1]
		if _, ok := interesting[current]; ok { // ignore if it's in the interesting list
			continue
		}
		if _, ok := uninteresting[current]; ok { // ignore if it's in the boring list
			continue
		}
		first := true
		ok := true
		over := false
		for i := 0; i < 8; i++ {
			x := current[0] + i%2
			y := current[1] + (x/2)%2
			z := current[2] + (x/4)%2
			gp := gridpoint{x, y, z}
			field := 0.

			if field, ok = grid[gp]; !ok {
				field = c.fn(c.gridToXYZ(gp))
				grid[gp] = field
			}
			if first {
				over = (field > c.level)
			}
			if (field > c.level) != over {
				interesting[current] = ccube{
					gp:    current,
					field: [2][2][2]float64{}, // fill later
				}
				break
			}
		}
		// TODO : disassemble this
	}
	return out
}
