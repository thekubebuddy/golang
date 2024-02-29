package utils

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Div(x, y int) float32 {
	if y == 0 {
		return -1
	} else {
		return float32(x / y)
	}
}
