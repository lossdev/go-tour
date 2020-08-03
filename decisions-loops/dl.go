package main

import (
    "fmt"
)

func main() {
    /*
        Decisions and Loops are some of the most basic bread and butter of not
        only a Go program, but programs in general. Without decisions and loops,
        programs would not be of much computational use (see Turing machines
        and their impact - https://en.wikipedia.org/wiki/Turing_machine).

        Decisions in go are done using conditional statements, and is something
        we all use in our daily lives. It's as simple as making decisions based
        on conditions - if it's hot outside, where a T-Shirt; if I'm hungry, I'll
        eat. Programs do the same thing, based on conditions you've set. The main
        two ways to achieve decisions are through if/else statements, and switch
        statements.

        Looping is the process of iterating multiple times based on a set condition,
        to the effect of manipulating other values or conditions. Looping in go is
        done exclusively through 'for' loops - there are no while / do while loops
        in go.
    */

    // first, let's set a boolean value
    a := false
    /* 
        now, let's introduce an if/else statement
        if statements do not need parentheses in Go, unlike some languages like
        C or Java. Also, else if / else statements MUST start after the preceding 
        closing brace }, not doing so will give a compilation error.
    */
    if a == true {
        fmt.Println("a is true")
    } else {
        fmt.Println("a is false")
    }

    // The above snippet can also be shortened to this (only for boolean values):
    if a {
        fmt.Println("a is true")
    } else {
        fmt.Println("a is false")
    }

    /*
        Switch statements are also a form of decision making, and are particularly
        useful when a target condition can be multiple values depending on other
        conditions. Each of those conditions you'd like to test for are called
        'cases', and can be chained one after another to cover the full range
        of values you want to test for. You can also set a default case,
        which will be executed when none of the previous cases are met.
    */

    // initialize a value
    b := 3
    // now, the switch statement
    switch b {
        case 1:
            fmt.Println("b is 1")
        case 2:
            fmt.Println("b is 2")
        case 3:
            fmt.Println("b is 3")
        default:
            fmt.Println("b is something else")
    }

    /*
        Now, let's try looping. Loops are identified with a 'for' statement. Similar
        to C and Java syntax, a for loop requires 3 different arguments, semicolon
        separated:

           - iterator variable
           - condition for loop end
           - iterator increment / decrement pattern

        The below loop will initialize a variable 'i' to start at 0, set its 
        break condition to be when i = 5 (ie the loop will keep going while
        i is less than 5), and that i will increment by one for every iteration
        of the loop. Iterators in a for loop can also decrement (like i--), and
        even skip by a multiple number. For example, i += 2 would increase i by
        2 every iteration of the loop instead of 1. for loops, like if statements,
        do not necessarily need enclosing parentheses.
    */

    // This loop will print "Hello!" 5 times
    for i := 0; i < 5; i++ {
        fmt.Println("Hello!")
    }

    /*
        To better demonstrate a switch statement's usage, let's now combine
        our knowledge of loops and switch statements:
    */
    // iterate from 1 to 4
    for i := 1; i <= 4; i++ {
        // set b equal to i
        b = i
        switch b {
            case 1:
                fmt.Println("b is 1")
            case 2:
                fmt.Println("b is 2")
            case 3:
                fmt.Println("b is 3")
            default:
                fmt.Println("b is something else")
        }
    }
    /*
        Now, we can start to see the use of the switch statement. b is changing
        its value on every iteration of the loop, and our switch statement can
        now pick up (some) of the different values b is being set to.
    */
}
