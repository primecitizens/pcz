// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package matrix

import (
	"unsafe"

	"github.com/primecitizens/std/core/num"
	"github.com/primecitizens/std/math"
)

// Mat4 is a 4x4 matrix in row major order.
//
// m[4*r + c] is the element in the r'th row and c'th column.
type Mat4[T num.Real] [16]T

func (m *Mat4[T]) SetIdentity() *Mat4[T] {
	*m = Mat4[T]{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
	return m
}

func (m *Mat4[T]) Translate(v Vec3[T]) *Mat4[T] {
	*(*[4]T)(unsafe.Pointer(&m[12])) = [4]T{
		m[0]*v[0] + m[4]*v[1] + m[8]*v[2] + m[12],
		m[1]*v[0] + m[5]*v[1] + m[9]*v[2] + m[13],
		m[2]*v[0] + m[6]*v[1] + m[10]*v[2] + m[14],
		m[3]*v[0] + m[7]*v[1] + m[11]*v[2] + m[15],
	}
	return m
}

func (m *Mat4[T]) Mul(x Mat4[T]) *Mat4[T] {
	*m = Mat4[T]{
		m[0]*x[0] + m[4]*x[1] + m[8]*x[2] + m[12]*x[3],
		m[1]*x[0] + m[5]*x[1] + m[9]*x[2] + m[13]*x[3],
		m[2]*x[0] + m[6]*x[1] + m[10]*x[2] + m[14]*x[3],
		m[3]*x[0] + m[7]*x[1] + m[11]*x[2] + m[15]*x[3],

		m[0]*x[4] + m[4]*x[5] + m[8]*x[6] + m[12]*x[7],
		m[1]*x[4] + m[5]*x[5] + m[9]*x[6] + m[13]*x[7],
		m[2]*x[4] + m[6]*x[5] + m[10]*x[6] + m[14]*x[7],
		m[3]*x[4] + m[7]*x[5] + m[11]*x[6] + m[15]*x[7],

		m[0]*x[8] + m[4]*x[9] + m[8]*x[10] + m[12]*x[11],
		m[1]*x[8] + m[5]*x[9] + m[9]*x[10] + m[13]*x[11],
		m[2]*x[8] + m[6]*x[9] + m[10]*x[10] + m[14]*x[11],
		m[3]*x[8] + m[7]*x[9] + m[11]*x[10] + m[15]*x[11],

		m[0]*x[12] + m[4]*x[13] + m[8]*x[14] + m[12]*x[15],
		m[1]*x[12] + m[5]*x[13] + m[9]*x[14] + m[13]*x[15],
		m[2]*x[12] + m[6]*x[13] + m[10]*x[14] + m[14]*x[15],
		m[3]*x[12] + m[7]*x[13] + m[11]*x[14] + m[15]*x[15],
	}
	return m
}

func (m Mat4[T]) MulV(x Mat4[T]) Mat4[T] {
	return Mat4[T]{
		m[0]*x[0] + m[4]*x[1] + m[8]*x[2] + m[12]*x[3],
		m[1]*x[0] + m[5]*x[1] + m[9]*x[2] + m[13]*x[3],
		m[2]*x[0] + m[6]*x[1] + m[10]*x[2] + m[14]*x[3],
		m[3]*x[0] + m[7]*x[1] + m[11]*x[2] + m[15]*x[3],

		m[0]*x[4] + m[4]*x[5] + m[8]*x[6] + m[12]*x[7],
		m[1]*x[4] + m[5]*x[5] + m[9]*x[6] + m[13]*x[7],
		m[2]*x[4] + m[6]*x[5] + m[10]*x[6] + m[14]*x[7],
		m[3]*x[4] + m[7]*x[5] + m[11]*x[6] + m[15]*x[7],

		m[0]*x[8] + m[4]*x[9] + m[8]*x[10] + m[12]*x[11],
		m[1]*x[8] + m[5]*x[9] + m[9]*x[10] + m[13]*x[11],
		m[2]*x[8] + m[6]*x[9] + m[10]*x[10] + m[14]*x[11],
		m[3]*x[8] + m[7]*x[9] + m[11]*x[10] + m[15]*x[11],

		m[0]*x[12] + m[4]*x[13] + m[8]*x[14] + m[12]*x[15],
		m[1]*x[12] + m[5]*x[13] + m[9]*x[14] + m[13]*x[15],
		m[2]*x[12] + m[6]*x[13] + m[10]*x[14] + m[14]*x[15],
		m[3]*x[12] + m[7]*x[13] + m[11]*x[14] + m[15]*x[15],
	}
}

func (m *Mat4[T]) MulScalar(x T) *Mat4[T] {
	*m = Mat4[T]{
		m[0] * x, m[1] * x, m[2] * x, m[3] * x,
		m[4] * x, m[5] * x, m[6] * x, m[7] * x,
		m[8] * x, m[9] * x, m[10] * x, m[11] * x,
		m[12] * x, m[13] * x, m[14] * x, m[15] * x,
	}
	return m
}

func (m *Mat4[T]) Rotate(rad float64, axis Vec3[T]) *Mat4[T] {
	x := float64(axis[0])
	y := float64(axis[1])
	z := float64(axis[2])
	n := math.Sqrt(x*x + y*y + z*z)
	x /= n
	y /= n
	z /= n

	xx := x * x
	yy := y * y
	zz := z * z
	c := math.Cos(rad)
	s := math.Sin(rad)
	oneMinusCosine := 1 - c

	r00 := xx + (1-xx)*c
	r01 := x*y*oneMinusCosine + z*s
	r02 := x*z*oneMinusCosine - y*s
	r10 := x*y*oneMinusCosine - z*s
	r11 := yy + (1-yy)*c
	r12 := y*z*oneMinusCosine + x*s
	r20 := x*z*oneMinusCosine + y*s
	r21 := y*z*oneMinusCosine - x*s
	r22 := zz + (1-zz)*c

	*(*[12]T)(unsafe.Pointer(m)) = [12]T{
		T(r00*float64(m[0]) + r01*float64(m[4]) + r02*float64(m[8])),
		T(r00*float64(m[1]) + r01*float64(m[5]) + r02*float64(m[9])),
		T(r00*float64(m[2]) + r01*float64(m[6]) + r02*float64(m[10])),
		T(r00*float64(m[3]) + r01*float64(m[7]) + r02*float64(m[11])),

		T(r10*float64(m[0]) + r11*float64(m[4]) + r12*float64(m[8])),
		T(r10*float64(m[1]) + r11*float64(m[5]) + r12*float64(m[9])),
		T(r10*float64(m[2]) + r11*float64(m[6]) + r12*float64(m[10])),
		T(r10*float64(m[3]) + r11*float64(m[7]) + r12*float64(m[11])),

		T(r20*float64(m[0]) + r21*float64(m[4]) + r22*float64(m[8])),
		T(r20*float64(m[1]) + r21*float64(m[5]) + r22*float64(m[9])),
		T(r20*float64(m[2]) + r21*float64(m[6]) + r22*float64(m[10])),
		T(r20*float64(m[3]) + r21*float64(m[7]) + r22*float64(m[11])),
	}

	return m
}

func (m *Mat4[T]) SetRotationX(rad float64) *Mat4[T] {
	c := math.Cos(rad)
	s := math.Sin(rad)

	*m = Mat4[T]{
		1, 0, 0, 0,
		0, T(c), T(s), 0,
		0, T(-s), T(c), 0,
		0, 0, 0, 1,
	}

	return m
}

func (m *Mat4[T]) RotateX(rad float64) *Mat4[T] {
	c := math.Cos(rad)
	s := math.Sin(rad)

	*(*[8]T)(unsafe.Pointer(&m[4])) = [8]T{
		T(c*float64(m[4]) + s*float64(m[8])),
		T(c*float64(m[5]) + s*float64(m[9])),
		T(c*float64(m[6]) + s*float64(m[10])),
		T(c*float64(m[7]) + s*float64(m[11])),

		T(c*float64(m[8]) - s*float64(m[4])),
		T(c*float64(m[9]) - s*float64(m[5])),
		T(c*float64(m[10]) - s*float64(m[6])),
		T(c*float64(m[11]) - s*float64(m[7])),
	}

	return m
}

func (m *Mat4[T]) SetRotationY(rad float64) *Mat4[T] {
	c := math.Cos(rad)
	s := math.Sin(rad)

	*m = Mat4[T]{
		T(c), 0, T(-s), 0,
		0, 1, 0, 0,
		T(s), 0, T(c), 0,
		0, 0, 0, 1,
	}

	return m
}

func (m *Mat4[T]) RotateY(rad float64) *Mat4[T] {
	c := math.Cos(rad)
	s := math.Sin(rad)

	m0 := *(*[4]T)(unsafe.Pointer(&m[0]))

	*(*[4]T)(unsafe.Pointer(&m[0])) = [4]T{
		T(c*float64(m0[0]) - s*float64(m[8])),
		T(c*float64(m0[1]) - s*float64(m[9])),
		T(c*float64(m0[2]) - s*float64(m[10])),
		T(c*float64(m0[3]) - s*float64(m[11])),
	}

	*(*[4]T)(unsafe.Pointer(&m[8])) = [4]T{
		T(c*float64(m[8]) + s*float64(m0[0])),
		T(c*float64(m[9]) + s*float64(m0[1])),
		T(c*float64(m[10]) + s*float64(m0[2])),
		T(c*float64(m[11]) + s*float64(m0[3])),
	}

	return m
}

func (m *Mat4[T]) SetRotationZ(rad float64) *Mat4[T] {
	c := math.Cos(rad)
	s := math.Sin(rad)

	*m = Mat4[T]{
		T(c), T(-s), 0, 0,
		T(s), T(c), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	return m
}

func (m *Mat4[T]) RotateZ(rad float64) *Mat4[T] {
	c := math.Cos(rad)
	s := math.Sin(rad)

	*(*[8]T)(unsafe.Pointer(m)) = [8]T{
		T(c*float64(m[0]) + s*float64(m[4])),
		T(c*float64(m[1]) + s*float64(m[5])),
		T(c*float64(m[2]) + s*float64(m[6])),
		T(c*float64(m[3]) + s*float64(m[7])),

		T(c*float64(m[4]) - s*float64(m[0])),
		T(c*float64(m[5]) - s*float64(m[1])),
		T(c*float64(m[6]) - s*float64(m[2])),
		T(c*float64(m[7]) - s*float64(m[3])),
	}

	return m
}

func SetPerspective4[T num.Int | num.Float](
	m *Mat4[T], fieldOfViewYInRadians, aspect, zNear, zFar float64,
) *Mat4[T] {
	var (
		f = math.Tan(math.Pi*0.5 - 0.5*fieldOfViewYInRadians)

		e10, e14 T
	)

	if math.IsInf(zFar, 0) {
		e10 = -1
		e14 = T(-zNear)
	} else {
		rangeInv := 1 / (zNear - zFar)
		e10 = T(zFar * rangeInv)
		e14 = T(zFar * zNear * rangeInv)
	}

	*m = Mat4[T]{
		T(f / aspect), 0, 0, 0,
		0, T(f), 0, 0,
		0, 0, e10, -1,
		0, 0, e14, 0,
	}

	return m
}
