package main

import (
    "fmt"
    "math"
    "strconv"
    "sync"
)

func main() {
    /*
        Maps are a data structure that holds key-value pairs. Keys need to be
        unique, but the same value can be assigned to multiple keys. maps need to
        have a defined key type and value type - they can be different between the
        two, but need to be consistent internally. For example, a valid map could
        have integer keys and string values, but all keys must be integers
        and all values must be strings. Let's define that map below using make():
    */
    m :=  make(map[int]string)
    m[1] = "foo"
    m[2] = "bar"
    m[4] = "baz"
    // maps support the len() function to find the amount of items in them
    l := len(m)
    fmt.Printf("%s: %d\n", "m len", l)
    // access map values
    fmt.Println(m[1])
    fmt.Println(m[2])
    // won't print anything
    fmt.Println(m[3])

    /*
        Goroutines are a pretty special Go implementation of concurrency and 
        multithreading - however, goroutines do not use as many resources as an
        extra thread or process. Goroutines behave like functions (and in fact
        DO execute functions), but do it in concert with the main execution 
        thread. The 'go' keyword signifies a goroutine, and the argument afterwards
        should be the function to start in the goroutine. Goroutines also allow
        for channels to be used.
    */

    // use a WaitGroup to wait for all goroutines to finish executing
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        // this is an anonymous function - it's not declared globally and doesn't
        // have a calling name. It's simply used here as a quick function not
        // intended to be reused
        go func(val int) {
            // close WorkGroup for this routine when done
            defer wg.Done()
            fmt.Println("Hello " + strconv.Itoa(val))
        }(i)
    }
    // wait here
    wg.Wait()
    fmt.Println("Hello 5")

    /*
        Channels are data flow agents that allow you to send and receive data. 
        Channels can be assigned to a variable and send and receive data in
        multiple functions or routines simultaneously, allowing for a good use
        in distributed workloads and messaging queues.
    */
    c := make(chan int)
    /*
        defer calls will execute when the function had ended execution - this is
        useful for things like open files, channels, etc that should be closed
        before the program ends.
    */
    defer close(c)
    /*
        Arrow notation (<-) denotes if data is being inserted or taken out of the
        channel, depending on what side the arguments are on.

        c <- var : send var into channel c
        var <-c  : receive from channel c and assign to var
    */
    go square(c, 4)
    // Receive value from channel and assign to v
    v := <-c
    fmt.Println(v)
}

func square(c chan int, val int) {
    c <- int(math.Pow(float64(val), 2.0))
}
