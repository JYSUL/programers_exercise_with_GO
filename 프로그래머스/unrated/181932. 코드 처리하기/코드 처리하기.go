func solution(code string) string {
    var mod int = 0
    var answer string 
    answer = ""
    for index, str := range code {
        if str == 49 {
            mod = (mod + 1) % 2
        } else {
            if index % 2 == mod {
                answer += string(str)
            }
        }
    }
    if answer == "" {
        answer = "EMPTY"
    }
    return answer
}