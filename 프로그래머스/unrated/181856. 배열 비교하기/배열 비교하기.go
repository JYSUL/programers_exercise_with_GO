func solution(arr1 []int, arr2 []int) int {
    var arr1_count, arr2_count int = 0, 0
    if len(arr1) != len(arr2) {
        arr1_count = len(arr1)
        arr2_count = len(arr2)
    } else {
        for i:=0; i< len(arr1); i++ {
            arr1_count += arr1[i]
            arr2_count += arr2[i]
        }
    }
    if arr1_count < arr2_count {
        return -1
    } else if arr1_count == arr2_count {
        return 0
    } else {
        return 1
    }
}