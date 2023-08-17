import "strings"
func solution(my_string string) []string {
    var answer_slice []string
    for _, str := range strings.Split(my_string, " ") {
        if str != "" {
            answer_slice = append(answer_slice, str)
        }
    }
    return answer_slice
}