func solution(brown int, yellow int) []int {
    all_tiles := brown + yellow
    for 세로 := 3; 세로 <= all_tiles / 2; 세로++ { //yellow > 1 -> 세로길이 3이상
        if all_tiles % 세로 != 0 {
            continue
        }
        가로 := all_tiles / 세로
        if (세로 - 2) * (가로 - 2) == yellow {
            return []int{가로, 세로}
        }
        
    }
    return []int{0, 0} // error
}