package main
 
import (
    "fmt"
    "math"
)
 
const N = 1000000

func primeNumber () []int64 {
    var x, y, n int64
    nsqrt := math.Sqrt(N)

    is_prime := [N]bool{}

    for x = 1; float64(x) <= nsqrt; x++ {
        for y = 1; float64(y) <= nsqrt; y++ {
            n = 4*(x*x) + y*y
            if n <= N && (n%12 == 1 || n%12 == 5) {
                is_prime[n] = !is_prime[n]
            }
            n = 3*(x*x) + y*y
            if n <= N && n%12 == 7 {
                is_prime[n] = !is_prime[n]
            }
            n = 3*(x*x) - y*y
            if x > y && n <= N && n%12 == 11 {
                is_prime[n] = !is_prime[n]
            }
        }
    }

    for n = 5; float64(n) <= nsqrt; n++ {
        if is_prime[n] {
            for y = n * n; y < N; y += n * n {
                is_prime[y] = false
            }
        }
    }

    is_prime[2] = true
    is_prime[3] = true

    primes := make([]int64, 0, N)
    for  x = 0; x < int64(len(is_prime))-1; x++ {
        if is_prime[x] {
            primes = append(primes, x)
        }
    }

    return primes
}

func main() {
    var number int64
    fmt.Println("Enter Number: ")
    fmt.Scan(&number)
    
    scanNumber := number

    if (number > 1000000000000) {
        fmt.Println("number must not more than 1,000,000,000,000")
    } else {
        fmt.Println("Proccessing, please wait")
        primes := primeNumber()
        //number = 600851475143 
        faktor := make([]int64,0,100000000)
        for number > 1 {
            for _, x := range primes {
                if (number % x) == 0 {
                    faktor = append(faktor, x)
                    number = number / x
                    break
                }
                if (x >= number) {
                    faktor = append(faktor, number)
                    break
                }
            }
        }
        fmt.Print("Prime Factor from " , scanNumber , " : ")
        fmt.Println(faktor)
    }
}