package contour

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
