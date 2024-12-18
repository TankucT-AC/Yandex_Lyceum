package main

import (
	"calculator/utils"
	"fmt"
)

func main() {
	data := "11-2+7-(3-5)"
	fmt.Println(utils.Calc(data))
}
