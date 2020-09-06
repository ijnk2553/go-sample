package main

import (
	"fmt"

	"go.sample.com/interface/controller"
)

func main() {
	out := "sss"
	cont := new(controller.Controller)
	output, err := cont.Run(out)
	if err != nil {
		fmt.Print(output)
	}

}
