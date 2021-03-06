/*
    Here, we declare our package as firstModule, not main. This means that this
    code cannot be directly ran, but the exposed methods can be called using 
    firstModule.Method()
*/

package firstModule

import (
    "fmt"
)

/*
    In go, functions that start with lowercase letters are unexported - meaning 
    that their visibility is only within the program they're currently in. Making
    a function visible to be called externally requires capitalizing it, as we'll 
    be doing with the SayHello method
*/

func SayHello() {
    fmt.Println("Hello from firstModule!")
}
