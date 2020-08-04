package main

import (
    "fmt"
    /*
        The strconv library contains some string casting methods that we'll be
        using later
    */
    "log"
    /*
        The log library is very commonly used in conjunction with the errors 
        library and error handling in general - log prints out some more 
        diagnostic information that fmt does not by default, making it useful
        for debugging and error messages
    */
    "strconv"
)

func main() {
    /*
        While our entrypoint is the main() function, we can also declare and use
        helper functions to accomplish the subtasks we need. Using functions helps
        you organize your code in a more readable format and reuse code to cut down
        on the total size of the program. Functions in just about any programming 
        language will share a few elements:

            - The function name
            - The function inputs
            - The function outputs (or return values)

        Just as we declared main using the 'func' keyword, other functions will
        also start with a 'func' keyword. Most times, functions will be globally
        scoped, however, go also supports inline (or function-within-function) 
        functioning, similar to Javascript. 
    */

    // Let's set up some variables
    a := 2
    // Now, if we wanted to print a, this would work:
    fmt.Println(a)
    // However, this would not work:
    // fmt.Println("a value: " + a)
    /*
        If you remove the comment and try and compile that, go will tell you that
        it cannot concatenate two mismatched types int and string. Unlike some 
        languages like Python or JavaScript, Go expects us to provide the same 
        type of variable in a print statement. To get around this issue, we'll
        have to perform a type-cast of the a variable and turn it into a string
        while printing it. This can be done with a call to strconv's Itoa() method
        (Integer TO Array - specifically a character array, or string).
    */
    fmt.Println("a value: " + strconv.Itoa(a))
    /*
        Now, let's try something similar using a function that takes in two numbers
        (as a string) and returns the result of multiplying them, as an int. We're
        also going to take advantage of Go's multiple returns to return two things:
        
            - Our answer
            - An error

        In Go, lots of functions will return the thing you want, as well as an error.
        If there is an error, the error variable will be set. If there isn't an 
        error, then the error variable wlll be nil - nil meaning the absence of
        a value at all, NOT zero. In our case, we will set both a value for the number
        we want, as well as an error if something goes wrong.
    */

    val, err := multiplyStrings("3", "4")
    /*
        Again, we'll want to check if error is NOT nil right after our function call.
        If it is not nil, then something went wrong, and we should print the error.
    */
    if err != nil {
        /*
            log's functions are similar to fmt's, except they also print some
            extra diagnostic information with the error we're printing.
        */
        log.Println(err)
    } else {
        fmt.Println(val)
    }
    /*
        Now, let's try and give multiplyStrings values that will not work. This
        time, the error will be printed.
    */
    val, err = multiplyStrings("xyz", "4")
    if err != nil {
        log.Println(err)
    } else {
        fmt.Println(val)
    }
    // you can also use the blank identifier _ in place of a return value -
    // useful for when you don't care about a certain value being returned, and 
    // don't want to set a variable to hold it
    _, err = multiplyStrings("2," "7")
    if err != nil {
        log.Println(err)
    }
}

/*
    Here is our new function multiplyStrings() - it takes in two string arguments,
    and returns two values - our integer value of multiplying the two strings
    together, and an error if something goes wrong.

    The return values MUST be enclosed by parentheses if there are multiple
    values being returned, like we have in this case. However, if there is only 
    one return value, then parentheses around it are not required.
*/

func multiplyStrings(num1 string, num2 string) (int, error) {
    // Atoi will return an integer if successful, and an error if not
    int1, err := strconv.Atoi(num1)
    // An integer type cannot be set to nil, so we'll just return 0 instead
    if err != nil {
        return 0, err
    }
    // Same test here
    int2, err := strconv.Atoi(num2)
    if err != nil {
        return 0, err
    }
    // return the product of the two numbers, and a nil (no) error
    return int1 * int2, nil
}
