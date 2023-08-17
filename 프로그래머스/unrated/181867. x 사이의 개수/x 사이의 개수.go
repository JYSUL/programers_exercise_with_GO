import "strings"
func solution(myString string) []int {
    var answer_slice []int 
    /*
    var start int = 0
    for index, alp := range myString {
        if string(alp) == "x" {
            answer_slice = append(answer_slice, index - start)
            start = index + 1
        }
    }
    answer_slice = append(answer_slice, len(myString) - start)
    */
    for _, str := range strings.Split(myString, "x") {
        answer_slice = append(answer_slice, len(str))
    }
    
    return answer_slice
    
}