//import "fmt"
import "strings"
func solution(s string) string {
    var answer string
    //문제 예제랑 채점결과랑 다르다.
    //all_str := strings.Fields(s)
    all_str := strings.Split(s, " ")
    
    for slice_index, str := range all_str {
        answer = ""
        for index, word := range str {
            if index % 2 == 0 {
                answer += strings.ToUpper(string(word))
            } else {
                answer += strings.ToLower(string(word))
            }
        }
        all_str[slice_index] = answer
    }
    return strings.Join(all_str," ")
}