package square

import (
	"math"
	"testing"
)

type testCase struct {
	sideLen float64
	sideNum SidesNumber
	want    float64
}

var results = []float64{706.8583}

func TestCalcSquare(t *testing.T) {

	sidesTestCases := []testCase{
		{15, SidesCircle, 706.8583},
		{20.35, SidesCircle, 1301.0042},
		{15, SidesTriangle, 97.4278},
		{15, SidesSquare, 225},
		{15, SidesNumber(5), 0},
	}

	for _, sidesTestCase := range sidesTestCases {
		got := CalcSquare(sidesTestCase.sideLen, sidesTestCase.sideNum)
		want := sidesTestCase.want

		if math.Abs(got-want) > 0.0001 {
			t.Errorf("got %v, wanted %v", got, want)
		} else {
			t.Logf("got %v", got)
		}
	}
}
