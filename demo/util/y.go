package util

const y int = 200 //unexported

//exported
func GetY() int {
	return y
}
