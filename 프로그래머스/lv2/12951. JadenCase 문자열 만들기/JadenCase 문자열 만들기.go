import ( 
    "strings"
)
func solution(s string) string {
    //strs := strings.Fields(s) fields 는 공백말고 다른 특수 기호들까지도 split한다.
    strs := strings.Split(s, " ")
    for index, str := range strs {
        strs[index] = strings.Title(strings.ToLower(str))
    }
    return strings.Join(strs," ")
}