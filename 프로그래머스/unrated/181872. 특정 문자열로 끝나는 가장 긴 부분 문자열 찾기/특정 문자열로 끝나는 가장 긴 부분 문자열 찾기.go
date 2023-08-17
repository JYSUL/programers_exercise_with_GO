//import "fmt"
func solution(myString string, pat string) string {
    for i := len(myString); i >= len(pat); i-- {
        if myString[i-len(pat):i] == pat {
            return myString[:i]
        }
    }
    return "No pat in myString"
}