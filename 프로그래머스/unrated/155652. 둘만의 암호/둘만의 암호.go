import (
    //"fmt"
)
func solution(s string, skip string, index int) string {
    var answer string
    skip_table := make(map[string]bool)
    for _, skipping := range skip {
        skip_table[string(skipping)] = true
    }
    for _, str := range s {
        count := 0
        for count < index {
            str++
            if str > 122 {
                str = 97
            }
            
            if !skip_table[string(str)] {
                count++
            }
        }
        answer += string(str)
    }
    return answer
}