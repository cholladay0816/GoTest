// This is a test package to learn how to write and import local modules

package mathpack

import (
	"math"
)

// Pow ...
func Pow(i float64, m float64) float64 {
	return math.Pow(i, m)
}
