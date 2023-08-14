func solution(my_string string, is_prefix string) int {
    if len(my_string) < len(is_prefix) {
        return 0
    }
    
    if my_string[:len(is_prefix)] == is_prefix {
        return 1
    } else {
        return 0
    }
}