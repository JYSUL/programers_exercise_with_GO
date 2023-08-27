import "fmt"
func solution(seoul []string) string {
    for index, name := range seoul {
        if name == "Kim" {
            return fmt.Sprintf("김서방은 %v에 있다", index)
        }
    }
    return "Not Found"
}