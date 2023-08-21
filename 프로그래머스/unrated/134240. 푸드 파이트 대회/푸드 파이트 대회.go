import "strconv"
func solution(food []int) string {
    var left, right string
    for index, num_of_food := range food {
        for i := 0; i < num_of_food / 2; i++ {
            left = left + strconv.Itoa(index)
            right = strconv.Itoa(index) + right
        }
    }
    return left + "0" + right
}