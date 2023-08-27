import (
    //"fmt"
    //"strings"
)
func solution(s string, n int) string {
    var answer string = ""
    for _, str := range s {
        if string(str) == " " {
            answer += " "
            continue
        }
        if str > 96 {
            str += rune(n)
            if str > 122 {
                str -= 26     // n < 25
            }       
        } else {
            str += rune(n)
            if str > 90 {
                str -= 26
            }
        }
        
        answer += string(str)
        
    }
    return answer
}


// azAZ 97 122 65 90