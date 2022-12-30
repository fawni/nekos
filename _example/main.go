// Licensed under the Open Software License version 3.0

package main

import (
	"fmt"
	"log"

	"github.com/fawni/nekos/v2"
)

func main() {
	owoified, err := nekos.Owoify("notices your bulge")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(owoified)

	catface, err := nekos.CatText()
	if err != nil {
		log.Fatalln(err)
	}
	eightBall, err := nekos.EightBall()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("8ball response: %s %s\n8ball image: %s", eightBall.Text, catface, eightBall.Image)
}
