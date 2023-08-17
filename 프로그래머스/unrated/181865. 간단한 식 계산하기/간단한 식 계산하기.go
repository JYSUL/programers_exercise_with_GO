import (
    "strconv"
    "strings"
)
func solution(binomial string) int {
    num1, err := strconv.Atoi(strings.Split(binomial," ")[0])
    if err != nil {
        panic("Wrong input")
    }
    operater := strings.Split(binomial," ")[1]
    num2, err := strconv.Atoi(strings.Split(binomial," ")[2])
    if err != nil {
        panic("Wrong input")
    }
    var answer int
    switch operater {
    case "+" :
        answer = num1 + num2
    case "-" :
        answer = num1 - num2
    case "*" :
        answer = num1 * num2        
    }
    return answer
}