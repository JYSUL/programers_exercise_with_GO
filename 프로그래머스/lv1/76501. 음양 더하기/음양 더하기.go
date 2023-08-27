func solution(absolutes []int, signs []bool) int {
    var result int = 0
    for i := 0; i < len(absolutes); i++ {
        if signs[i] {
            result += absolutes[i]
        } else {
            result -= absolutes[i]
        }
    }
    return result
}