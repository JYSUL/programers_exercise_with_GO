//import "fmt"
func solution(x int, y int, n int) int {
    var checking_slice [][]int = [][]int{[]int{x,0}}
    var index int = 0
    var checked map[int]bool = make(map[int]bool)
    
    for index < len(checking_slice) {
        num, new_count := checking_slice[index][0], checking_slice[index][1]
        index++
        if checked[num] {
            continue
        }
        checked[num] = true
        if num == y {
            return new_count
        }
        for _, new_num := range []int{num+n, num*2, num*3} {
            if checked[new_num] {
                continue
            }
            if new_num <= y {
                checking_slice = append(checking_slice, []int{new_num, new_count + 1})
            }
        }
    }
    
    //fmt.Println(checking_slice, index)
    return -1
}