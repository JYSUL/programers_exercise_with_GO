import(
    "strings"
)
func solution(my_string string, indices []int) string {
    var answer string = ""
    list := strings.Split(my_string, "")
    
    for _, index := range indices {
        list[index] = "A"
    }
    
    for _, str := range list {
        if str != "A" {
            answer += str
        }
    }
    return answer
}