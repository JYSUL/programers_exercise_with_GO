import "strings"
func solution(my_string string, alp string) string {
    return strings.Replace(my_string, alp, strings.ToUpper(alp), -1)
}