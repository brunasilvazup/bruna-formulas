// This is the formula implementation class.
// Where you will code your methods and manipulate the inputs to perform the specific operation you wish to automate.

package formula

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

// Formula is the struct aboud inputs
type Formula struct {
	Name, BirthDate, Country, State string
}

// Run is the main function of formula
func (f Formula) Run(writer io.Writer) {
	var result string

	result += fmt.Sprintf("Hello World!\n")
	result += color.FgGreen.Render(fmt.Sprintf("My name is %s.\n", f.Name))
	result += color.FgYellow.Render(fmt.Sprintf("My birthday is %s.\n", f.BirthDate))
	result += color.FgBlue.Render(fmt.Sprintf("My country is %s.\n", f.Country))
	result += color.FgGreen.Render(fmt.Sprintf("My state is %s.\n", f.State))

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
