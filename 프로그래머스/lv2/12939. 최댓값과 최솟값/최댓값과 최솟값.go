import (
    "strings"
    "strconv"
)
func solution(s string) string {
    var min, max int
    nums := strings.Split(s, " ")
    min, _ = strconv.Atoi(nums[0])
    max, _ = strconv.Atoi(nums[0])
    
    
    
    for _, str := range strings.Split(s, " ") {
        num, err := strconv.Atoi(str)
        if err != nil {
            panic("panic")
        }
        if num < min {
            min = num
        }
        if num > max {
            max = num
        }
        
    }
    
    return strconv.Itoa(min) + " " + strconv.Itoa(max)
}