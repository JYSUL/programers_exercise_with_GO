//import "fmt"
func solution(keymap []string, targets []string) []int {
    var answer_slice []int
    key_table := make(map[string]int) 
    for _, keys := range keymap {
        for index, key := range keys {
            if key_table[string(key)] == 0 {
                key_table[string(key)] = index + 1
            } else if key_table[string(key)] > index + 1{
                key_table[string(key)] = index + 1
            }
             
        }
    }
    //fmt.Println(key_table)
    for _, target := range targets {
        answer := 0
        for _, alp := range target {
            if key_table[string(alp)] == 0 {
                answer = -1
                break
            } else {
                answer += key_table[string(alp)]
            }
        }
        answer_slice = append(answer_slice, answer)
    }
    return answer_slice
}