package godemo

import (
	"github.com/sprufu/godemo/div"
	"github.com/sprufu/godemo/test"
)

func SumInt(a, b int) int {
	return a + b
}

func SumFloat64(a, b float64) float64 {
	return a + b
}

func Div(a, b float64) float64 {
	return div.DivFloat(a, b)
}

func MyPrint(a interface{}) {
	test.Print(a)
}