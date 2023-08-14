//import "fmt"
func solution(number string) int {
    var answer int = 0
    for _, num := range number {
        answer += int(num - 48)
        answer %= 9
    }
    return answer
}