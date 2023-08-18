import (
    //"fmt"
    "regexp"
)
func solution(myStr string) []string {
    var answer []string
    for _, str := range regexp.MustCompile("[abc]").Split(myStr, -1) {
        if len(str) != 0 {
            answer = append(answer,str)
        }
    }
    if len(answer) == 0 {
        answer = []string{"EMPTY"}
    }
    return answer
}