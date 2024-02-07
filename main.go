package main

import (
    "flag"
    "fmt"
    "github.com/automoto/diff-check/pkg"
)

func generateAnswerString(decimalPlace uint, outputNum float64) string {
    switch decimalPlace {
    case 0:
        return fmt.Sprintf("%.0f", outputNum)
    case 1:
        return fmt.Sprintf("%.1f", outputNum)
    case 2:
        return fmt.Sprintf("%.2f", outputNum)
    case 3:
        return fmt.Sprintf("%.3f", outputNum)
    case 4:
        return fmt.Sprintf("%.4f", outputNum)
    }
    return fmt.Sprintf("Decimal places from 0 to 4 are supported for the diff-check cli, your option %d is not supported", decimalPlace)
}

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
    flag.UintVar(&decimalPlaces, "decimal-places", 2, "How many decimal places would you like to round your answer to? Values from 0-4 supported")
    flag.UintVar(&decimalPlaces, "dp", 2, "How many decimal places would you like to round your answer to?")
    flag.StringVar(&typeCalc, "type", "diff", "What type of calculation do you want to do? valid options are 'change' or 'diff'")
    flag.StringVar(&typeCalc, "t", "diff", "What type of calculation do you want to do, Percent Difference or Percent Change? valid options are 'change' or 'diff'")
    flag.Parse()
    answer, err = pkg.Start(numOne, numTwo, typeCalc, decimalPlaces)
    if err != nil {
        fmt.Println(err)
    }
    s := generateAnswerString(decimalPlaces, answer)
    fmt.Println(s)
}
