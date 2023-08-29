import ( 
    "strings"
)
func solution(s string) string {
    //strs := strings.Fields(s)
    strs := strings.Split(s, " ")
    for index, str := range strs {
        strs[index] = strings.Title(strings.ToLower(str))
    }
    return strings.Join(strs," ")
}