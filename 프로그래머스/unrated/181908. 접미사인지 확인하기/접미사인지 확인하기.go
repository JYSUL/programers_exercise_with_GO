func solution(my_string string, is_suffix string) int {
    if len(my_string) < len(is_suffix) {
        return 0
    }
    
    if my_string[len(my_string) - len(is_suffix):] == is_suffix {
        return 1
    } else {
        return 0
    }
}