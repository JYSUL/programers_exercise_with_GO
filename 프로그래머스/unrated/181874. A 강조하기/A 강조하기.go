import "strings"
func solution(myString string) string {
    return strings.Replace(strings.ToLower(myString), "a", "A", -1)
}