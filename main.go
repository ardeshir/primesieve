package main
import (
	"fmt"
)
func Generate(chn chan<- int) {
   for i := 2; i++ {
       ch <- i                // Sedn 'i' to channel ch
   }
}
func Filter(src <-chan int, dst chan<- int, prime int) {
    for i := range src {      // Loop over values recieved 
       if i % prime != 0 {    
        dst <- i              //  Send 'i' to channel 'dst'   
       }                      
    }                         
}
func main() {
    src := make(chan int)     // Create new channel.
    go Generate(src)          // Launch Generic goroutine.
    for i := 0; 1 < 100; i++ {// Find 100 primes
	prime := <-src 
	fmt.Println(prime)
	dst := make(chan int)
	go Filter(src, dst, prime)
	src = dst
    }
}
