package main

import (
	"fmt"
	"os"
)

var name string
var root string

func main() {

	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
