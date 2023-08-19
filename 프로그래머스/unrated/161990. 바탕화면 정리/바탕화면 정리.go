func solution(wallpaper []string) []int {
    var lux, luy, rdx, rdy int = len(wallpaper[0]), len(wallpaper), 0, 0
    for now_y, cals := range wallpaper {
        for now_x, file := range cals {
            if string(file) == "#" {
                if now_x < lux {
                    lux = now_x
                }
                if now_x > rdx {
                    rdx = now_x
                }
                
                if now_y < luy {
                    luy = now_y
                }
                if now_y > rdy {
                    rdy = now_y
                }
            }
        }
    }
    return []int{luy, lux, rdy + 1, rdx + 1}
}