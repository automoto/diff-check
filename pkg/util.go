package pkg

import (
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
    return "", fmt.Errorf("invalid calculation type: %s It needs to be 'change' or 'diff", l)
}

func Start(firstNum, secondNum float64, typeCalc string, decimalPlaces uint) (float64, error) {
    var answer float64
    t, err := parseType(typeCalc)
    if err != nil {
        return 0, err
    }
    if t == "change" {
        answer, err = difference.PercentChange(firstNum, secondNum, decimalPlaces)
    }
    if t == "diff" {
        answer, err = difference.PercentDifference(firstNum, secondNum, decimalPlaces)
    }
    if err != nil {
        return 0, err
    }
    return answer, nil
}
