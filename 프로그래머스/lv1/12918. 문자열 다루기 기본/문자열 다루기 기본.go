func solution(s string) bool {
    check_table := make(map[rune]bool)
    for _, numeric := range "1234567890" {
        check_table[numeric] = true
    }
    
    if len(s) != 4 && len(s) != 6 {
        return false
    }
    
    for _, str := range s {
        if !check_table[str] {
            return false
        }
    }
    return true
}