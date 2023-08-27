import (
    "strings"
    "sort"
)
func solution(s string) string {
    splited := strings.Split(s, "")
    sort.Sort(sort.Reverse(sort.StringSlice(splited)))
    return strings.Join(splited,"")
}