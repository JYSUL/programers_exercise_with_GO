import (
    "math"
    //"fmt"
)
func solution(numbers int, limit int, power int) int {
    var answer int = 0
    
    for number := 1; number <= numbers; number++ {
        count := 0
        num := math.Sqrt(float64(number))
        if num == math.Floor(num) {
            count = -1
        }
        for i := 1; i <= int(num); i++ {
            if number % i == 0 {
                count += 2
            }
        }
        
        if count > limit {
            answer += power
        } else {
            answer += count
        }
        //fmt.Println(answer, count)
        
    }
    return answer
}