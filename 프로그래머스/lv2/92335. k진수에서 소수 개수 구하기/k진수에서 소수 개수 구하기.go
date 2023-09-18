//import "fmt"
import (
    "strconv"
    "strings"
    "math"
)

func solution(n int, k int) int {
    var count = 0
    
    //n을 k진수로 바꾸고 0으로 split -> primes에 저장
    str := strconv.FormatInt(int64(n), k)
    str = strings.Replace(str, "0", " ", -1)
    primes := strings.Fields(str)
    
    //primes에 있는 숫자들이 소수이면 count++
    for _, prime := range primes {
        num, err := strconv.Atoi(prime)
        if err != nil {
            panic("strconv.AToi line14")
        }
        
        if isPrime(num) {
            count++
        }
        
    }
    
    return count
}

//소수판별 함수 int(sqrt(n)) + 1 까지
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    num := float64(n)
    num = math.Sqrt(num)
    number := int(num)
    for i := 2; i < number + 1; i++ {
        if n % i == 0 {
            return false
        }
    }
    
    return true
}