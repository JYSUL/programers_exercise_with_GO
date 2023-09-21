//import "fmt"

import (
    "strings"
    "strconv"
    "sort"
)

//fee_cal calcluate fee by in, out time
func fee_cal(fees []int, time int) int {
    var result int = fees[1]
    var total_time int = time - fees[0]
    if total_time > 0 {
        temp := total_time / fees[2]
        if total_time % fees[2] != 0 {
            temp++
        }
        result += temp * fees[3]
    }
    return result
}

//func main
func solution(fees []int, records []string) []int {
    var result []int
    var key_map map[string]int = make(map[string]int)
    var index int = 0
    type car struct {
        number      string
        time_in     int
        total_time  int
    }
    var data []car
    
    
    for _, record := range records {        
        //number = car_number, direction = in or out, total_time = hour * 60 + minute
        record_array := strings.Split(record, " ")
        number, direction := record_array[1], record_array[2]
        hour, err := strconv.Atoi(record_array[0][:2])
        if err != nil {
            panic("strconv Atoi")
        }
        minute, err := strconv.Atoi(record_array[0][3:])
        if err != nil {
            panic("strconv Atoi")
        }
        total_time := hour * 60 + minute
        
        
        //차 번호에 맞게 data에 struct형태로 defalt저장. 해당 순서에 맞는 index를 key_map에 저장 "car_num" -> index
        car_index, flag := key_map[number]
        if !flag { // flag == false, 해당 차번호가 최초입력이면
            key_map[number], car_index = index, index
            data = append(data, car{number : number, time_in : 0, total_time :0})
            index++
        }
        
        //입력이 in이면 23:59에 출차로 시간계산을 해서 defalt로 저장
        //입력이 out이면 23:59출차에 대한 계산을 무효하고, out시간에 맞게 계산해서 저장
        switch direction {
        case "IN" :
            data[car_index].time_in = total_time
            data[car_index].total_time += 1439 - total_time
        case "OUT" :
            data[car_index].total_time -= 1439 - data[car_index].time_in
            data[car_index].total_time += total_time - data[car_index].time_in
        }
    }
    
    //차 번호를 사전순에 맞게 출력하기위해 key_map에서 차번호들의 slice를 keys에 저장후 정렬
    keys := []string{}
    for key, _ := range key_map {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    
    for _, key := range keys {
        result = append(result, fee_cal(fees, data[key_map[key]].total_time))
    }
    
    return result
}