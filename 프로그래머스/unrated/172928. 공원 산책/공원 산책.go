import (
    "strconv"
    //"fmt"
)
func solution(park []string, routes []string) []int {
    var start_y, start_x int
    var flag bool
    for y, low := range park {
        for x, cal := range low {
            if string(cal) == "S" {
                start_y, start_x = y, x
            }
        }
    }
    for _, route := range routes {
        direction := string(route[0])
        distance, err := strconv.Atoi(string(route[2]))
        if err != nil {
            panic("input error, distance of route")
        }
        flag = true
        
        switch direction {
        case "N" :
            if start_y - distance >= 0 {
                for i := 0; i <= distance; i++ {
                    if string(park[start_y-i][start_x]) == "X" {
                        flag = false
                        break
                    }
                }
                if flag {
                    start_y -= distance
                }
            }
        case "S" :
            if start_y + distance < len(park) {
                for i := 0; i <= distance; i++ {
                    if string(park[start_y+i][start_x]) == "X" {
                        flag = false
                        break
                    }
                }
                if flag {
                    start_y += distance
                }
            }
        case "W" :
            if start_x - distance >= 0 {
                for i := 0; i <= distance; i++ {
                    if string(park[start_y][start_x-i]) == "X" {
                        flag = false
                        break
                    }
                }
                if flag {
                    start_x -= distance
                }
            }            
        case "E" :
            if start_x + distance < len(park[0]) {
                for i := 0; i <= distance; i++ {
                    if string(park[start_y][start_x+i]) == "X" {
                        flag = false
                        break
                    }
                }
                if flag {
                    start_x += distance
                }
            }
        }
        //fmt.Println(start_y,start_x)    
    }
    return []int{start_y,start_x}
}