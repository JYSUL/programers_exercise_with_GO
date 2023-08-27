import (
    //"fmt"
    "sort"
)
func solution(n int, lost []int, reserve []int) int {
    sort.Ints(lost)
    var lost_else_mine []int
    reserve_table := make(map[int]bool)
    for _, res := range reserve {
        reserve_table[res] = true
    }
    
    //reserve가 있는사람이 도난 당했으면 자기자신이 입는다.
    for _, student := range lost {
        if reserve_table[student] {
            reserve_table[student] = false
        } else {
            lost_else_mine = append(lost_else_mine, student)
        }
    }
    
    
    for _, student := range lost_else_mine {
        if reserve_table[student - 1] {
            reserve_table[student - 1] = false
        } else if reserve_table[student + 1] {
            reserve_table[student + 1] = false
        } else {
            n--
        }
    }
    
    return n
}