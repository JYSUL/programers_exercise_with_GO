//import "fmt"
func solution(s string) int {
    var answer int = 0
    var same_num, diff_num int = 0, 0
    var str byte = s[0]
    for i := 0; i < len(s) - 1; i++ { // 마지막은 분리가 되든, 안되는 +1하기위해서 i < len(s) - 1
        if str == s[i] {
            same_num++
        } else {
            diff_num++
        }
        
        if same_num == diff_num {
            same_num, diff_num = 0, 0
            str = s[i+1]            
            answer++
        }
    }
    //fmt.Println(string(str))
    //마지막
    answer++ 
    return answer
}