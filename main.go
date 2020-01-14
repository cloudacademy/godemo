package main

import (
	"fmt"

	"github.com/cloudacademy/godemo/math"
	"github.com/cloudacademy/godemo/util"
)

func main() {
	x := util.GetX()
	y := util.GetY()

	fmt.Println(x)
	fmt.Println(y)

	sum := math.Add(x, y)
	uuid := math.GetUuid()

	fmt.Println(sum)
	fmt.Println(uuid)
}
