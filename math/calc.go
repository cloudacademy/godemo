package math

import "github.com/google/uuid"

//exported
func Add(num1 int, num2 int) int {
	return num1 + num2
}

func GetUuid() uuid.UUID {
	return uuid.New()
}
