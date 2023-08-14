import (
    //"fmt"
    "strings"
)
func solution(myString string, pat string) int {
    myString = strings.ToLower(myString)
    pat = strings.ToLower(pat)
    for i := 0; i < len(myString) - len(pat) + 1; i++ {
        
        if myString[i:i+len(pat)] == pat {
            return 1
        }
    }
    return 0
}