import (
    //"fmt"
)

func isPerfect(s string) bool {
    table := map[string]string{")":"(", "]":"[", "}":"{"}
    var result string = ""
    for _, str := range s {
        if len(result) == 0 {
            result += string(str)
            continue
        } 
        if table[string(str)] == string(result[len(result)-1]) {
            result = result [:len(result)-1]
        } else {
            result += string(str)
        }
        
    }
    if len(result) == 0 {
        return true
    } else {
        return false
    }
}


func solution(s string) int {
    var answer int = 0
    for i := 0; i< len(s); i++ {
        if isPerfect(s[i:] + s[:i]) {
            answer++
        }
    }
    return answer
}