func solution(my_string string, overwrite_string string, s int) string {
    var s1 int = len(overwrite_string) + s
    return my_string[:s] + overwrite_string + my_string[s1:]
}