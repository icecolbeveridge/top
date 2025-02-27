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
// Gridwise, I think the vertices of the cube are at (0,0,0) to (6,6,6)
// in the obvious sense.
// The edges of length 1 are at (3,[0|6], [0|6]) and permutations thereof.
// The sqrt(2) edges are at (3,3,[0|6]) and perms thereof.
// I'll need to think about the directions.
// The faces are... where are they? The orthogonal ones are centred at
// places like (0,2,2) -- the edges for that are (0,3,0), (0,0,3) and (0,3,3),
// which gives us an average.
// What about the slanted faces? Sample points are (0,0,0), (6,6,0), and (0,6,6), giving (2,4,2).
// And the tets? The one with points (0,0,0), (0,6,0), (6,6,0) and (0,0,6) has centroid... (1.5, 3, 1.5).
// Drat, we'll need to double it all (unless I can be happy with a fudge.)
// The big central tet is at (3,3,3).
