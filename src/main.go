package main

import (
	"ether-convertor/toolkit"
	"fmt"
	"os"
)

func main() {
	to, from, value,help := toolkit.ParseFlags()
	if help == "help" {
		toolkit.PrintHelp()
		os.Exit(0)
	}
	if check := toolkit.CheckInputs(to,from,value); !check {
		fmt.Printf("Exiting....")
		os.Exit(1)
	}
	fmt.Println(to, from)
	fmt.Printf("%s\n", value)
}


