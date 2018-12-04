package gocv-funcs

import (
	"image"

	"gocv.io/x/gocv"
)

func ApplyWatermark(src, wm gocv.Mat, dx, dy int) gocv.Mat {
	new := src.Clone()

	offset := image.Point{
		X: dx,
		Y: new.Rows() - dy - wm.Rows(),
	}

	for y := 0; y < wm.Rows(); y++ {
		for x := 0; x < wm.Cols(); x++ {
			wmDot := GetVecbAt(wm, y, x)
			srcDot := GetVecbAt(new, offset.Y+y, offset.X+x)
			for c := 0; c < new.Channels(); c++ {
				switch wmDot[3] {
				case 0:
					continue
				case 255:
					srcDot[c] = wmDot[c]
				default:
					alpha := float64(wmDot[3]) / 255.0
					srcDot[c] = uint8(float64(wmDot[c])*alpha + float64(srcDot[c])*(1-alpha))
				}
			}
			srcDot.SetVecbAt(new, offset.Y+y, offset.X+x)
		}
	}
	return new
}

type Vecb []uint8

func GetVecbAt(m gocv.Mat, row int, col int) Vecb {
	ch := m.Channels()
	v := make(Vecb, ch)

	for c := 0; c < ch; c++ {
		v[c] = m.GetUCharAt(row, col*ch+c)
	}

	return v
}

func (v Vecb) SetVecbAt(m gocv.Mat, row int, col int) {
	ch := m.Channels()

	for c := 0; c < ch; c++ {
		m.SetUCharAt(row, col*ch+c, v[c])
	}
}
