func solution(a int, b int, c int, d int) int {
    var answer int = 0
    var table map[int]int
    table = make(map[int]int)
    for _, i := range []int{a,b,c,d} {
        table[i] = 0
    }
    for _, i := range []int{a,b,c,d} {
        table[i] += 1
    }
    
    switch len(table) {
    case 1 :
        answer = a * 1111
    case 2 :
        answer = 0
        if table[a] == 1 || table[a] == 3 {
            for num, count := range table {
                if count == 3 {
                    answer += 10 * num
                } else {
                    answer += num
                }
            }
            answer *= answer
        } else {
            nums := []int{}
            for num, _ := range table {
                nums = append(nums, num)
            }
            answer = (nums[0] + nums[1]) * (nums[0] - nums[1])
            if answer < 0 {
                answer *= -1
            }
        }
    case 3:
        answer = 1
        for num, count := range table {
            if count == 1 {
                answer *= num
            }
        }
    case 4:
        answer = 7
        for num, _ := range table {
            if answer > num {
                answer = num
            }
        }
    }
    
    return answer
}