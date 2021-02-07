package toolkit

import (
	"flag"
	"fmt"
	"strconv"
)

// ParseFlags :  parse and return input flags
func ParseFlags() (string, string, string, string) {
	fmt.Println("Welcome to ether convertor")
	to := flag.String("to", "wei", "Unit of value to convert to")
	from := flag.String("from", "wei", "Unit of value to convert from")
	value := flag.String("value", "", "Enter the value in this field")
	help := flag.String("help", "printing help", "Print all the options present")
	flag.Parse()

	return *to, *from, *value, *help
}

// PrintHelp : print all the help guide to the user
func PrintHelp() {
	fmt.Println("Printing help")
}

// CheckInputs : check all input for validity
func CheckInputs(to, from, value string) bool {
	if !valueCheck(to) {
		fmt.Println("Incorrect `to` value")
		return false
	}
	if !valueCheck(from) {
		fmt.Println("Incorrect `from` value")
		return false
	}

	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Problem parsing `value`")
		return false
	}

	return true
}

func valueCheck(val string) bool {
	if val == "wei" || val == "kwei" || val == "mwei" || val == "gwei" || val == "microether" || val == "milliether" || val == "ether" {
		return true
	} else {
		return false
	}
}
