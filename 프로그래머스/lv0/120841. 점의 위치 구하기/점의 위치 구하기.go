func solution(dot []int) int {
    var x int = dot[0]
    var y int = dot[1]
    var answer int = 0
    if x > 0 {
        if y > 0 {
            answer = 1
        } else {
            answer = 4
        }
    } else {
        if y > 0 {
            answer = 2
        } else {
            answer = 3
        }
    }
    return answer
}