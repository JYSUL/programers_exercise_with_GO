func solution(todo_list []string, finished []bool) []string {
    var answer []string
    for index, flag := range finished {
        if !flag {
            answer = append(answer, todo_list[index])
        }
    }
    return answer
}