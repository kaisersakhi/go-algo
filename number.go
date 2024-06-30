package main

import (
    "strconv"
    "strings"
)

import "go-algo/shared"

type Number struct {
    Value int
}

func (p Number) Compare(q Number) int {
    if p.Value > q.Value {
        return shared.GreaterThan
    } else if p.Value < q.Value {
        return shared.LessThan
    }
    return shared.EqualTo
}

func NumberArray(ints []int) []Number{
    intsLength := len(ints)
    numbers := make([]Number, intsLength)
    
    for i := 0; i < intsLength; i++ {
        numbers[i] = New(ints[i])
    }
    
    return numbers 
}

func New(intN int) Number{
    return Number{Value: intN}
}

func ToString(numbers []Number) string{
    numbersLength := len(numbers)
    
    var outputBuilder strings.Builder
    outputBuilder.WriteString("[")
    
    for i := 0; i < numbersLength; i ++ {
        if (numbersLength - 1) == i {
            outputBuilder.WriteString(strconv.Itoa(numbers[i].Value))
        }else {
            outputBuilder.WriteString(strconv.Itoa(numbers[i].Value) + ", ")
        }
    }
    
    outputBuilder.WriteString("]")
    return outputBuilder.String()
}