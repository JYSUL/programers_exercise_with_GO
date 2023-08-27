import (
    //"fmt"
    "strconv"
    "strings"
    "sort"
)
func solution(n int64) int64 {
    str_num := strconv.FormatInt(n, 10)
    num_splited := strings.Split(str_num,"")
    sort.Sort(sort.Reverse(sort.StringSlice(num_splited)))
    numeric_num, err := strconv.ParseInt(strings.Join(num_splited,""), 10, 64)
    if err != nil {
        panic("asdf")
    }
    return numeric_num
}