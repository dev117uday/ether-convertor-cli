package convertor

func wei(to, value *string) float64 {
	num := parseFloat(*value)
	switch *to {
	case "wei":
		return num
	case "kwei":
		return num * 1000
	case "mwei":
		return num * 1000 * 1000
	case "gwei":
		return num * 1000 * 1000 * 1000
	case "microether":
		return num * 1000 * 1000 * 1000 * 1000
	case "milliether":
		return num * 1000 * 1000 * 1000 * 1000 * 1000
	case "ether":
		return num * 1000 * 1000 * 1000 * 1000 * 1000
	}
	return num
}

func kwei(to, value *string) float64 {
	return 1
}

func mwei(to, value *string) float64 {
	return 1
}

func gwei(to, value *string) float64 {
	return 1
}

func microether(to, value *string) float64 {
	return 1
}

func milliether(to, value *string) float64 {
	return 1
}

func ether(to, value *string) float64 {
	return 1
}


