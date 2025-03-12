package fieldline

import "top/top"

// Let's implement 3D RK4 with adaptive stepsize
// This is pretty much a direct port from NR

func rk4(y, dydx top.Vector,
	x, h float64,
	derivs func(x float64, y top.Vector) top.Vector,
) top.Vector {

	hh := 0.5 * h
	h6 := h / 6.
	xh := x + hh
	yt := y.Add(dydx.ScalarMult(hh))
	dyt := derivs(xh, yt)
	yt = y.Add(dyt.ScalarMult(hh))
	dym := derivs(xh, yt)
	yt = y.Add(dym.ScalarMult(h))
	dym = dym.Add(dyt)
	dyt = derivs(x+h, yt)
	yy := dym.ScalarMult(2.).Add(dyt).Add(dydx)
	return y.Add(yy.ScalarMult(h6))
}
