import "strconv"
func solution(s string) []int {
    count_zero := 0
    count_step := 0
    
    for s != "1" {
        new_s := ""
        for _, str := range s{
            if string(str) == "0" {
                count_zero++
            } else {
                new_s += string(str)
            }
        }
        s = strconv.FormatInt(int64(len(new_s)), 2)
        count_step++
    }
    
    return []int{count_step, count_zero}
}