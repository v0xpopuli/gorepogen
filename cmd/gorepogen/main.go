package main

import (
	"fmt"
	"os"
)

func main() {

	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("\n%v", err)
	}

}
