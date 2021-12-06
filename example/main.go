// Licensed under the Open Software License version 3.0

package main

import (
	"fmt"

	"github.com/x6r/nekos"
)

func main() {
	fmt.Println(nekos.Owoify("notices your bulge"))
	fmt.Println(nekos.Why())
	fmt.Println(nekos.CatText())
	fmt.Println(nekos.Cuddle())
	fmt.Println(nekos.Kitsune())
	eightBall := nekos.EightBall()
	fmt.Printf("8ball response: %s\n8ball image: %s", eightBall.Text, eightBall.Image)
}
