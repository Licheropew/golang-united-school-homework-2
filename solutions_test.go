package square

import (
	"math"
	"testing"
)

const (
	SidesTriangles Sides = 3
	SidesSquares   Sides = 4
	SidesCircles   Sides = 0
	SidesOthers    Sides = 6
)

func TestCalcSquare(t *testing.T) {
	var x float64 = 10.0
	var result float64
	t.Run("Triangle", func(t *testing.T) {
		funcResult := CalcSquare(x, SidesTriangles)
		result = 10 * 10 * 0.25 * math.Sqrt(3)
		if funcResult != result {
			t.Errorf("Real result not equal expected result. %f != %f\n", funcResult, result)
		}
	})
	t.Run("Square", func(t *testing.T) {
		funcResult := CalcSquare(x, SidesSquares)
		result = 10 * 10
		if funcResult != result {
			t.Errorf("Real result not equal expected result. %f != %f\n", funcResult, result)
		}
	})
	t.Run("Circle", func(t *testing.T) {
		funcResult := CalcSquare(x, SidesCircles)
		result = math.Pi * 100
		if funcResult != result {
			t.Errorf("Real result not equal expected result. %f != %f\n", funcResult, result)
		}
	})
	t.Run("Other", func(t *testing.T) {
		funcResult := CalcSquare(x, SidesOthers)
		result = 0
		if funcResult != result {
			t.Errorf("Real result not equal expected result. %f != %f\n", funcResult, result)
		}
	})
}
