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
            if len(sub_belt) == 0 {
                result = index
                break
            }
            
            if sub_belt[len(sub_belt)-1] == order {
                sub_belt = sub_belt[:len(sub_belt)-1]
            } else {
                result = index
                break
            }
        }
    } 
    
    return result
}

/*
스택과 남아있는 택배의 관계를 생각하면
order가 순차적인 입력이 아니면 중단한다.
탑승될 택배가 해당 범위의 +1 or -1인 경우에만 추가적인 진행을 한다.
*/