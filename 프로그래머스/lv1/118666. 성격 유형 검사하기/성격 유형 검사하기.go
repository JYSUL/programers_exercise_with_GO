//import "fmt"
func solution(survey []string, choices []int) string {
    var result string
    points_table := make(map[string]int)
    for index, choice := range survey {
        switch string(choice[0]) {
        case "R" :
            points_table["R"] += (choices[index] - 4) * -1
        case "T" :
            points_table["R"] += (choices[index] - 4)
        case "C" :
            points_table["C"] += (choices[index] - 4) * -1
        case "F" :
            points_table["C"] += (choices[index] - 4)
        case "J" : 
            points_table["J"] += (choices[index] - 4) * -1
        case "M" :
            points_table["J"] += (choices[index] - 4)
        case "A" :
            points_table["A"] += (choices[index] - 4) * -1
        case "N" :
            points_table["A"] += (choices[index] - 4)
        }
    }
    //fmt.Println(points_table)
    
    if points_table["R"] >= 0 {
        result += "R"
    } else {
        result += "T"
    } 
    if points_table["C"] >= 0 {
        result += "C"
    } else {
        result += "F"
    } 
    if points_table["J"] >= 0 {
        result += "J"
    } else {
        result += "M"
    } 
    if points_table["A"] >= 0 {
        result += "A"
    } else {
        result += "N"
    } 
    
    
    
    
    
    return result
}

/*
R C J A
T F M N



*/