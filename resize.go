package gocv-funcs

import (
	"image"
	"math"

	"gocv.io/x/gocv"
)

func ResizeWithFill(src gocv.Mat, wantx, wanty int, fill gocv.Scalar) gocv.Mat {
	resizedImg := gocv.NewMat()

	multiplier := math.Min(float64(wantx)/float64(src.Cols()), float64(wanty)/float64(src.Rows()))
	width := int(math.Round(float64(src.Cols()) * multiplier))
	height := int(math.Round(float64(src.Rows()) * multiplier))

	gocv.Resize(
		src,
		&resizedImg,
		image.Point{X: width, Y: height},
		0,
		0,
		gocv.InterpolationLinear,
	)

	if resizedImg.Cols() != wantx || resizedImg.Rows() != wanty {
		filled := gocv.NewMatWithSize(wanty, wantx, src.Type())

		dx := (wantx - width) / 2
		dy := (wanty - height) / 2

		modx := dx*2 - (wantx - width)
		mody := dy*2 - (wanty - height)

		filled.SetTo(fill)

		roi := filled.Region(image.Rectangle{
			Min: image.Point{X: dx, Y: dy},
			Max: image.Point{X: filled.Cols() - dx + modx, Y: filled.Rows() - dy + mody},
		})

		// Because we lost pointer to roi underlying data after CopyTo call
		// Possible memory leak here
		m := roi
		defer m.Close()

		resizedImg.CopyTo(&roi)
		resizedImg.Close()

		return filled
	}
	return resizedImg
}
