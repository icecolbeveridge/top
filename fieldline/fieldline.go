package fieldline

import "top/top"

// Let's implement 3D RK4 with adaptive stepsize
// This is pretty much a direct port from NR

type derivsFunc func(float64, top.Vector) top.Vector

func rk4(y, dydx top.Vector,
	x, h float64,
	derivs derivsFunc,
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

func rk_simple(vstart top.Vector, x1, x2 float64, nstep int, derivs derivsFunc) []top.Vector {
	y := make([]top.Vector, nstep+1)
	y[0] = vstart
	v := vstart
	xx := make([]float64, nstep+1)
	xx[0] = x1
	x := x1
	h := (x2 - x1) / float64(nstep)
	for k := 1; k <= nstep; k++ {
		dv := derivs(x, v)
		v = rk4(v, dv, x, h, derivs)
		if x+h == x {
			panic("Insignificant stepsize")
		}
		x += h
		xx[k] = x
		y[k] = v
	}
	return y
}
