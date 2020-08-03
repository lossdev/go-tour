package main

/*
    Go supports the common variable types, such as integers, floats, string,
    boolean, etc. Go supports both implicit variable assignments (weak typing),
    as well as explicit variable assignments (strong typing). Implicit assignments
    are achieved through using the := identifier like this:

        a := 2

    Explicit typing can also be done using the var keyword. Unlike most other 
    languages, Go expects the name of the variable to precede the type of that
    variable. For example, you could type the above line as this as well:

        var a int
        a = 2

        OR

        var a = 2

    While this is obviously longer than the := paradigm, there are certain cases
    where you will need to use the var syntax over the := syntax, most notably,
    when declaring a global variable. The := syntax cannot be used to assign
    a variable while in the global scope; instead, the var syntax must be used.
*/

import (
    "fmt"
)


// Let's initialize a couple of global variables here

var a = 2
var b = 4

func main() {
    // Boolean values can be set with 'true' or 'false' values
    c := true
    fmt.Println(c) // true
    // First, let's print our a and b values
    fmt.Println(a) // 2
    fmt.Println(b) // 4
    // We can modify a or b's value easily
    a += 1         // set a equal to itself plus 1
    fmt.Println(a) // 3
    d := 1.23      // float64 value
    fmt.Println(d) // 1.23
    e := "Hello!"  // string type
    fmt.Println(e) // "Hello!"
    /*
        When a variable this is in scope has already been set, reassigning it to
        another value is done using just =, not :=. := is ONLY used for 
        initialization AND assignment of a variable. For example, since a is
        already in scope, we can set it equal to something else:
    */
    a = 4
    fmt.Println(a)
}
