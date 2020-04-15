package util

import (
	"fmt"
	"math"
	"strconv"
)

var (
	sizeUnits = []string{"KB", "MB", "GB", "TB"}
)

func FormatSize(size int64) string {
	if size < 1024 {
		return strconv.FormatInt(size, 10)
	}
	fsize := float64(size)
	for i, unit := range sizeUnits {
		min, max := math.Pow(1024.0, float64(i+1)), math.Pow(1024.0, float64(i+2))
		if fsize < max {
			fsize /= min
			return fmt.Sprintf("%.2f %s", fsize, unit)
		}
	}
	return strconv.FormatInt(size, 10)
}
