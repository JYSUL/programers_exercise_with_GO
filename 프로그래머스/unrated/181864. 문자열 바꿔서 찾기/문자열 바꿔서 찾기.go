func solution(myString string, pat string) int {
    var new_pat string
    for _, alp := range pat {
        if string(alp) == "A" {
            new_pat += "B"
        } else {
            new_pat += "A"
        }
    }
    
    for i := 0; i < len(myString) - len(new_pat) + 1; i++ {
        if myString[i:i+len(new_pat)] == new_pat {
            return 1
        }
    }
    
    
    return 0
}