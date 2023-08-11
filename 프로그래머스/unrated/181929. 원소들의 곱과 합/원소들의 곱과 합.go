func solution(num_list []int) int {
    var sum_of_slice int = 0
    var product_of_slice int = 1
    for _, num := range num_list {
        sum_of_slice += num
        product_of_slice *= num
    }
    if sum_of_slice * sum_of_slice > product_of_slice {
        return 1
    } else {
        return 0
    }
}