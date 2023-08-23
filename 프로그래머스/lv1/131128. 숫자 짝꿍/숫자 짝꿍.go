import "strings"
//import "fmt"
func solution(X string, Y string) string {
    var num_string string = "9876543210"
    var answer string
    //var count int
    for _, num := range num_string {        
        if strings.Count(X, string(num)) > strings.Count(Y, string(num)) {
            answer += strings.Repeat(string(num), strings.Count(Y, string(num)))
        } else {
            answer += strings.Repeat(string(num), strings.Count(X, string(num)))
        }
    }
    
    if len(answer) == 0 {
        return "-1"
    } else if strings.HasPrefix(answer, "0") {
        return "0"
    } else {
        return answer
    }
}