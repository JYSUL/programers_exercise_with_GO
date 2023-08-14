import (
    //"fmt"
    "strconv"
)

func solution(intStrs []string, k int, s int, l int) []int {
    var answer []int
    for _, str := range intStrs {
        new_num, err := strconv.Atoi(str[s:s+l])
        if err != nil {
            continue
        }
        if new_num > k {
            answer = append(answer, new_num)
        }
    }
    return answer
}