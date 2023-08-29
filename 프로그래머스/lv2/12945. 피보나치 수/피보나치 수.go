func solution(n int) int {
    temp := 1
    pivo := 0
    
    for i := 0; i < n; i++ {
        pivo, temp = (pivo + temp) % 1234567 , pivo
    }
    return pivo
}