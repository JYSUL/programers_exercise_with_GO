//import "fmt"
func solution(answers []int) []int {
    first_student  := []int{1, 2, 3, 4, 5}
    second_student := []int{2, 1, 2, 3, 2, 4, 2, 5}
    third_student  := []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}
    count_table := make(map[int]int)
    var result []int
    var max_count int = 0
    
    
    for index, answer := range answers {
        if answer == first_student[index%5] {
            count_table[1]++
        }
        if answer == second_student[index%8] {
            count_table[2]++
        }
        if answer == third_student[index%10] {
            count_table[3]++
        }
    }
    
    for i := 1; i < 4; i++ {
        if count_table[i] > max_count {
            max_count = count_table[i]
        }
    }
    
    for i := 1; i< 4; i++ {
        if count_table[i] == max_count {
            result = append(result, i)
        }
    }
    
    
    return result
}