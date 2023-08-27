import (
    //"fmt"
    //"sort"
    "strings"
    "strconv"
)
func solution(n int64) []int {    
    var num_splited []int    
    for _, num := range strings.Split(strconv.FormatInt(n,10), "") {
        int_num, err := strconv.Atoi(string(num))
        if err != nil {
            panic("panic")
        }
        num_splited = append(num_splited, int_num)
    }
    for i, j := 0, len(num_splited)-1; i < j; i, j = i+1, j-1 {
        num_splited[i], num_splited[j] = num_splited[j], num_splited[i]
    }
    return num_splited
}