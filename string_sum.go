package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var (
		oneNumber, twoNumber string
		one, two, count      int
	)
	input = strings.ReplaceAll(input, " ", "")
	if input == "" {
		return "", fmt.Errorf("nothing to calculate. %w", errorEmptyInput)
	}
	for i, v := range input {
		V := string(v)
		if string(V) == "+" && count < 2 {
			count++
			oneNumber = input[:i]
			twoNumber = input[i+1:]
		} else if count == 2 {
			return "", fmt.Errorf("the number of operands is more than two. %w", errorNotTwoOperands)
		} else if count < 1 && i+1 == len(input) {
			return "", fmt.Errorf("number of operands must be equal two. %w", errorNotTwoOperands)
		} else if count == 0 && i+1 == len(input) {
			return "", fmt.Errorf("nothing to calculate. %w", errorEmptyInput)
		}
	}
	one, err = strconv.Atoi(oneNumber)
	if err != nil {
		fmt.Println(err)
	}
	two, err = strconv.Atoi(twoNumber)
	if err != nil {
		fmt.Println(err)
	}
	return strconv.FormatInt(int64(one+two), 10), nil
}
