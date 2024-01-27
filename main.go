package main

import (
    "flag"
    "github.com/automoto/diff-check/pkg"
)

func main() {
    var numOne float64
    var numTwo float64
    var decimalPlaces int
    var typeCalc string
    flag.Float64Var(&numOne, "first", 0, "first number to do calculations on")
    flag.Float64Var(&numOne, "f", 0, "first number to do calculations on")
    flag.Float64Var(&numTwo, "second", 0, "second number to do calculations on")
    flag.Float64Var(&numTwo, "s", 0, "second number to do calculations on")
    flag.IntVar(&decimalPlaces, "decimal-places", 2, "How many decimal places woudl you like to round your answer to?")
    flag.IntVar(&decimalPlaces, "dp", 2, "How many decimal places woudl you like to round your answer to?")
    flag.StringVar(&typeCalc, "type", "diff", "What type of calculation do you want to do? valid options are 'change' or 'diff'")
    flag.StringVar(&typeCalc, "type", "diff", "What type of calculation do you want to do, Percent Difference or Percent Change? valid options are 'change' or 'diff'")
    pkg.Start(numOne, numTwo, typeCalc, decimalPlaces)
}
