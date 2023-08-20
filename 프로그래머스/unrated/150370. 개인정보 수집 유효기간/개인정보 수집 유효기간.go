import (
    //"fmt"
    "strconv"
    "strings"
)
const Day_Of_Year int = 28*12
const Day_Of_Month int = 28

func dateToNum (date string) int {
    dates := strings.Split(date,".") 
    year, _ := strconv.Atoi(dates[0])
    month, _ := strconv.Atoi(dates[1])
    day, _ := strconv.Atoi(dates[2])
    return year * Day_Of_Year + month * Day_Of_Month + day
}


func solution(today string, terms []string, privacies []string) []int {
    term_table := make(map[string]int)    
    date_of_today := dateToNum(today)
    var answer_slice []int
    
    for _, term := range terms {
        month, _ := strconv.Atoi(strings.Split(term," ")[1])
        term_table[strings.Split(term," ")[0]] = month
    }
    
    for index, privacy := range privacies {
        priv := strings.Split(privacy," ")
        //fmt.Println(date_of_today, dateToNum(priv[0]), term_table[priv[1]])
        if date_of_today >= dateToNum(priv[0]) + term_table[priv[1]] * Day_Of_Month {
            answer_slice = append(answer_slice, index+1)
        }
    }
    return answer_slice
}