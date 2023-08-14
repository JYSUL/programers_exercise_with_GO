//import "fmt"
func solution(my_string string, queries [][]int) string {
    var new_string string = ""
    for _, query := range queries {
        new_string = ""
        start, end := query[0], query[1]
        new_string += my_string[:start]
        for i := end; i > start-1; i-- {
            new_string += string(my_string[i])
        }
        new_string += my_string[end+1:]
        my_string = new_string
    }
    return my_string
}