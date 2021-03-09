// This is the main class.
// Where you will extract the inputs asked on the config.json file and call the formula's method(s).

package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	input1 := os.Getenv("INPUT_NAME")
	input2 := os.Getenv("INPUT_BIRTH_DATE")
	input3 := os.Getenv("INPUT_COUNTRY")
	input4 := os.Getenv("INPUT_STATE")

	formula.Formula{
		Name:      input1,
		BirthDate: input2,
		Country:   input3,
		State:     input4,
	}.Run(os.Stdout)
}
