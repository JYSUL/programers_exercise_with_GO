import (
    //"fmt"
    "time"
    "strings"
)
func solution(a int, b int) string {
    //days_of_week := []string{"SUN","MON","TUE","WED","THU","FRI","SAT"}
    //layout := "2016-01-01"
    
    t := time.Date(2016, time.Month(a), b, 0, 0, 0, 0, time.UTC)
    //fmt.Printf("%v, %T", t.Weekday().String(), t.Weekday().String() )
    return strings.ToUpper(t.Weekday().String())[:3]
}