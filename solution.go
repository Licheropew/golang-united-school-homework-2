package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
	SidesCircle   Sides = 0
	SidesOther    Sides = 6
)

type Sides int

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	if sidesNum == SidesTriangle {
		return math.Pow(sideLen, 2) * math.Sqrt(3) * 0.25
	}
	if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}
	if sidesNum == SidesCircle {
		return math.Pow(sideLen, 2) * math.Pi
	}
	return 0
}
