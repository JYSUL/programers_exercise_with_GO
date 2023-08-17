import "strings"
func solution(strArr []string) []string {
    for index, str := range strArr {
        if index % 2 == 0 {
            strArr[index] = strings.ToLower(str)
        } else {
            strArr[index] = strings.ToUpper(str)
        }
    }
    return strArr
}