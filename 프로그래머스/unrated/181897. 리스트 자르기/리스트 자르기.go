func solution(n int, slicer []int, num_list []int) []int {
    var start, end, dis int = 0, len(num_list), 1
    var answer []int
    
    switch n {
    case 1 :
        end = slicer[1] + 1
    case 2 :        
        start = slicer[0]
    case 3 :
        start, end = slicer[0], slicer[1] + 1
    case 4 :
        start, end, dis = slicer[0], slicer[1] + 1, slicer[2]
    }
    
    for i := start; i < end; i += dis {
        answer = append(answer, num_list[i])
    }
    return answer
}