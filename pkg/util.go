package pkg

import (
    "errors"
    "fmt"
    "github.com/sydneyshile/go-difference-check"
    "strings"
)

func parseType(typeCalc string) (string, error) {
    s := strings.TrimSpace(typeCalc)
    l := strings.ToLower(s)
    if (l == "change") || (l == "diff") {
        return l, nil
    }
    return "", errors.New(fmt.Sprintf("Invalid calculation type: %s It needs to be 'change' or 'diff", l))
}

func Start(firstNum, secondNum float64, typeCalc string, decimalPlaces int) (float64, error) {
    var answer float64
    t, err := parseType(typeCalc)
    if err != nil {
        return "", err
    }
    if t == "change" {
        answer, err = difference.PercentChange(firstNum, secondNum, decimalPlaces)
    }
    if t == "diff" {
        answer, err = difference.PercentDifference(firstNum, secondNum, decimalPlaces)
    }
    return answer, nil
}
