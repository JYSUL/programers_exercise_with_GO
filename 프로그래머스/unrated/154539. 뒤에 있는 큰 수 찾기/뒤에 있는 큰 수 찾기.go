//import "fmt"

func solution(numbers []int) []int {
    var queue []int
    var result []int
    
    //numbers 뒤부터 확인한다.
    for i := len(numbers)-1; i >= 0; i-- {  
        
        //numbers의 뒤부터 queue를 만드는데, 만약 확인하는 값보다 작은값은 사이에서 확인할 필요가 없으므로 삭제한다.
        for len(queue) >0 &&queue[len(queue)-1] <= numbers[i] {
            queue = queue[:len(queue)-1]            
        }
        
        //queue의 길이가 0이면 내 뒤로 큰수가 없다는 뜻이므로 -1, 그 외에는 queue[-1]을 result에 append한다.
        //그냥 append하고 reverse하는 방법도 있고, 애초에 앞에 넣는 방법도 있다. 앞에 append 하는 방법은 너무 느리다
        //테스트11     4.65ms vs 6550.21ms
        if len(queue) == 0 {
            result = append(result, -1)
            //result = append([]int{-1}, result...)
        } else {
            result = append(result, queue[len(queue)-1])
            //result = append([]int{queue[len(queue)-1]}, result...)            
        }
        queue = append(queue, numbers[i])
    }
    //reverse in memory
    for i, j := 0, len(result)-1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    } 
    
    return result
}