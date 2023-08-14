func solution(str_list []string) []string {
    for index, str := range str_list {
        if str == "r" {
            return str_list[index+1:]
        } else if str == "l" {
            return str_list[:index]
        }
    }
    return []string{}
}