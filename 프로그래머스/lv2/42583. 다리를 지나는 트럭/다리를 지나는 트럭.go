//import "fmt"
func solution(bridge_length int, weight int, truck_weights []int) int {
    var bridge []int = make([]int, bridge_length)
    var bridge_index, truck_index, sum_of_weight int
    var count int
    
    //실제 트럭 시뮬레이터를 돌려본다.
    //bridge[bridge_index:] + bridge[:bridge_index] 가 실제 다리고 tick마다 bridge_index++을 한다.
    for truck_index < len(truck_weights) {
        sum_of_weight -= bridge[bridge_index]
        bridge[bridge_index] = 0
        
        if sum_of_weight + truck_weights[truck_index] <= weight {
            bridge[bridge_index] = truck_weights[truck_index]
            sum_of_weight += truck_weights[truck_index]
            truck_index++
        } 
        
        count++
        bridge_index = (bridge_index + 1) % len(bridge)
        
        //fmt.Println(bridge)
    }
    
    
    
    return count + len(bridge)
}