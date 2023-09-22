//import "fmt"

func solution(topping []int) int {
    var left_check map[int]bool = make(map[int]bool)
    var right_check map[int]int = make(map[int]int)
    var left_count, right_count int
    var result int
    
    //right checking
    for _, num := range topping {
        if right_check[num] == 0 {
            right_count++
        }
        right_check[num]++
    }
    
    for _, num := range topping {
        right_check[num]--
        if right_check[num] == 0 {
            right_count--
        }
        if left_check[num] == false {
            left_check[num] = true
            left_count++
        }
        //fmt.Println(right_count, left_count)
        if right_count == left_count {
            result++
        }
        
    }
    
    return result
}