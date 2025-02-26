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
