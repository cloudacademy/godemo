package util

const x int = 100 //unexported

//exported
func GetX() int {
	return x
}
