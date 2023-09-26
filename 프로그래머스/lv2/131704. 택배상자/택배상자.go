//import "fmt"
func solution(orders []int) int {
    var result int = len(orders)
    var sub_belt []int
    var cargo int = 1
    
    for index, order := range orders {
        if order > cargo {
            for ; cargo < order; cargo++ {
                sub_belt = append(sub_belt, cargo)
            }
            cargo++
        } else if order == cargo {
            cargo++
        } else {
            if len(sub_belt) == 0 || sub_belt[len(sub_belt)-1] != order {
                result = index
                break
            }            
            sub_belt = sub_belt[:len(sub_belt)-1]            
        }
    } 
    
    return result
}
