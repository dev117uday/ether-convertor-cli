package convertor

import (
	"strconv"
)

func ConvertorReceiver(to, from, value *string) float64 {
	switch *from {
	case "wei":
		return wei(to, value)
	case "kwei":
		return kwei(to, value)
	case "mwei":
		return mwei(to, value)
	case "gwei":
		return gwei(to, value)
	case "microether":
		return microether(to, value)
	case "milliether":
		return milliether(to, value)
	case "ether":
		return ether(to, value)
	}
	return 1
}

func parseFloat(value string) float64 {
	val,err := strconv.ParseFloat(value,64)
	if err != nil {
		panic("Something went wrong")
	}
	return val
}
