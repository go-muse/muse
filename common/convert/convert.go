package convert

import (
	"fmt"
)

// AddUint8Int8 returns sum of two uint8 and int8 numbers.
func AddUint8Int8(a uint8, b int8) uint8 {
	res := int16(a) + int16(b)

	if res < 0 || res > 255 {
		panic(fmt.Sprintf("AddUint8Int8: Sum '%d' is out of uint8 range", res))
	}

	return uint8(res)
}

// SubUint8Uint8 returns difference between two uint8 numbers.
func SubUint8Uint8(a, b uint8) int8 {
	var diff int

	if a >= b {
		diff = int(a) - int(b)
	} else {
		diff = -1 * (int(b) - int(a))
	}

	if diff < -128 || diff > 127 {
		panic(fmt.Sprintf("SubUint8Uint8: Difference %d is out of int8 range", diff))
	}

	return int8(diff)
}
