package main
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for num := range jobs {
        fmt.Printf("Worker %d processing %d\n", id, num)
    }
}

func main() {
    jobs := make(chan int)
    var wg sync.WaitGroup

    
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(i, jobs, &wg)
    }

   
    for i := 1; i <= 20; i++ {
        jobs <- i
    }
    close(jobs)

    wg.Wait()
    fmt.Println("All jobs done!")
}
