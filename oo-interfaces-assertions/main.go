package main

import (
    "example/queue"
    "fmt"
    "log"
)

func main() {
    q := queue.NewQueue()
    q.Add(1)
    q.Add(2)
    q.Add(3)
    val, err := q.Remove()
    if err != nil {
        log.Println(err)
    } else {
        // a type assertion is required when interface{} values are used. Since
        // we used int types, we can use an int type assertion
        fmt.Println(val.(int))
    }
    val, err = q.Remove()
    if err != nil {
        log.Println(err)
    } else {
        fmt.Println(val.(int))
    }
    val, err = q.Remove()
    if err != nil {
        log.Println(err)
    } else {
        fmt.Println(val.(int))
    }
    // this next call should error, and return our error we created
    val, err = q.Remove()
    if err != nil {
        log.Println(err)
    } else {
        fmt.Println(val.(int))
    }
}
