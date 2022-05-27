package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MyError struct {
	Message string
}

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands       = errors.New("expecting two operands, but received more or less")
	myErr               error = MyError{Message: "не соотвествует"}
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

func (e MyError) Error() string {
	return e.Message
}

func StringSum(input string) (output string, err error) {
	var outSum, count int
	var ic, inputClear string
	inputnew := strings.Split(input, " ")
	for _, p := range inputnew {
		ic += p
	}
	if ic == "" {
		return "NO", errorEmptyInput
	} else {
		for _, v := range ic {
			V := string(v)
			if strings.ContainsAny("0123456789+-", V) == true {
				inputnew := strings.Split(V, "+")
				for _, p := range inputnew {
					inputClear += p
					if V != "0" {
						count++
						if count > 2 {
							return "NO", errorNotTwoOperands
						}
					}
				}
			} else {

			}
		}
		for i := 0; i < len(inputClear); i++ {
			W := string(inputClear[i])
			if W == "-" {
				out, err := strconv.Atoi(inputClear[i : i+2])
				i++
				if err != nil {
					fmt.Println(err)
				}
				outSum += out
			} else {
				out, err := strconv.Atoi(W)
				if err != nil {
					fmt.Println(err)
				}
				outSum += out
			}
		}
		output = strconv.FormatInt(int64(outSum), 10)
		return output, nil
	}
}
