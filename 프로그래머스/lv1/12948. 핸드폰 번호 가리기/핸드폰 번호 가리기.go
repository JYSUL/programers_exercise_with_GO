func solution(phone_number string) string {
    var result string
    for i := 0; i < len(phone_number) - 4; i++ {
        result += "*"
    }
    result += phone_number[len(phone_number)-4:]
    return result
}