package main

import (
    "flag"
    "fmt"
    "github.com/automoto/diff-check/pkg"
)

func main() {
    var numOne float64
    var numTwo float64
    var decimalPlaces uint
    var typeCalc string
    var answer float64
    var err error

    flag.Float64Var(&numOne, "first", 0, "first number to do calculations on")
    flag.Float64Var(&numOne, "1", 0, "first number to do calculations on")
    flag.Float64Var(&numTwo, "second", 0, "second number to do calculations on")
    flag.Float64Var(&numTwo, "2", 0, "second number to do calculations on")
    flag.UintVar(&decimalPlaces, "decimal-places", 2, "How many decimal places would you like to round your answer to?")
    flag.UintVar(&decimalPlaces, "dp", 2, "How many decimal places would you like to round your answer to?")
    flag.StringVar(&typeCalc, "type", "diff", "What type of calculation do you want to do? valid options are 'change' or 'diff'")
    flag.StringVar(&typeCalc, "t", "diff", "What type of calculation do you want to do, Percent Difference or Percent Change? valid options are 'change' or 'diff'")
    flag.Parse()
    answer, err = pkg.Start(numOne, numTwo, typeCalc, decimalPlaces)
    s := fmt.Sprintf("%e", answer)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(s)
}
