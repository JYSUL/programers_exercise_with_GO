//import "fmt"


func solution(priorities []int, location int) int {
    var result int
    var index int = 0
    var priority_counts map[int]int = make(map[int]int)
    
    //더 확인 해야하는 우선순위가 있는지 확인하기위한 count 테이블
    for _, priority := range priorities {
        priority_counts[priority]++
    }
    
    //9 부터 시작해서 우선순위가 남아있으면 해당 index까지 탐색해서 index를 order에 입력
    checker := 9    
    for checker > 0 {
        for true {
            //더이상 탐색할 우선순위가 없으면 다음 우선순위를 탐색
            if priority_counts[checker] == 0 {
                checker--
                break
            }
            if priorities[index] == checker {
                
                //찾고자하는 location 순서가 되면 이때까지의 순서+1 return
                if index == location {
                    return result + 1
                }
                result++
                priority_counts[checker]--
            }
            //index는 순환으로 계속해서 돈다
            index = (index + 1) % len(priorities)
        }
    }
    //오류면 -1
    return -1
}