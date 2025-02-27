package main

import (
	"fmt"
	"golang-language/capitulo-2/pacotes/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CTof(tempconv.BoilinC))
}
