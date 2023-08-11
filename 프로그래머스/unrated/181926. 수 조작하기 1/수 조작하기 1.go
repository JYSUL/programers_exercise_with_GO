func solution(n int, control string) int {
    for _, input := range control {        
        switch string(input) {
        case "w" :
            n += 1
        case "s" :
            n -= 1
        case "d" :
            n += 10
        case "a" :
            n -= 10
        }
        
    }
    return n
}