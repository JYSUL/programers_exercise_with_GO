func solution(want []string, number []int, discount []string) int {
    want_table := map[string]int{}
    discount_table := map[string]int{}
    var answer int = 0
    var flag bool = true
    
    //want와 number로 map을 만들고, discount의 10개의 map과 비교해서 같으면 count++
    for index, wan := range want {
        want_table[wan] = number[index]
    }
    for i := 0; i < 10; i++ {
        discount_table[discount[i]]++
    }
    
    //제일처음이나 제일 마지막둘중 하나의 확인
    flag = true
    for _, wan := range want {
        if want_table[wan] != discount_table[wan] {
            flag = false
            break
        }
    }
    if flag {
        answer++
    }
    
    for i:=0; i< len(discount)-10; i++ {
        discount_table[discount[i]]--
        discount_table[discount[i+10]]++
        
        flag = true
        for _, wan := range want {
            if want_table[wan] != discount_table[wan] {
                flag = false
                break
            }
        }
        if flag {
            answer++
        }
    }
    
    
    
    
    return answer
}