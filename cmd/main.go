package main

import (
	"fmt"
	"os"
)

func main() {

	app := createApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
