func solution(numLog []int) string {
    var diff int
    var command string = ""
    for index, num := range numLog[1:] {
        diff = num - numLog[index]
        switch diff {
        case 1 :
            command += "w"
        case -1:
            command += "s"
        case 10 :
            command += "d"
        case -10 :
            command += "a"
        }
    }
    return command
}