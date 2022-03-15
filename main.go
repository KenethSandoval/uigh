package main

import (
	"fmt"
	"os"

	"github.com/KenethSandoval/uigh/ui"
)

func main() {
	if err := ui.NewProgram("stiv").Start(); err != nil {
		fmt.Println("Could not start uigh", err)
		os.Exit(1)
	}
}
