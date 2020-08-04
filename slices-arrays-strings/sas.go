package main

import (
    "fmt"
    "strings"
)

func main() {
    /*
        Arrays and slices are similar, but subtly different. However, they both
        seek to accomplish the same thing - containing a row of the same data
        type but differing contents, which can be accessed later at 
        any time. Where they are different - elements can be appended to and taken
        off of a slice, whereas array sizes are immutable - once they've been 
        allocated, they cannot be extended or deleted. When arrays are declared,
        a size must be specified to them; otherwise, it's a slice.

        I personally don't ever use or have a need for arrays. Not only are they
        inflexible, they also can cause some headaches when passing them to and 
        from functions - for example, passing an array to a function passes a 
        COPY of the array, not the array itself. Therefore, the original array
        is never modified, something that I've tripped up with personally.

        Because of those issues and more (such as arrays not being interchangeable
        with slices, and therefore not able to be passed in as arguments requiring
        a slice in functions), I would recommend staying away from them and just 
        learning slices and their methods.
    */

    // declaring an array:
    var arr = [4]int{1, 3, 5, 2}
    // arr is now locked to 4 items. The elements themselves can be modified, but
    // the amount of elements must stay the same
    fmt.Println(arr)
    // this, however, is the same array expressed as a slice:
    var sl = []int{1, 3, 5, 2}
    fmt.Println(sl)
    // we can also call the make() builtin function to create a slice of initial
    // values of a particular size
    sl2 := make([]int, 5)
    fmt.Println(sl2)
    // removing elements is a bit tricker than adding - if we wanted to remove the 
    // 2nd element in sl, we would add up the elements up to the second, then after
    // the second. we use the _ operator to signify that we don't care about the
    // return value and don't want to set a variable to hold it
    sl = append(sl[:2], sl[3:]...)
    fmt.Println(sl)
    // now, let's add something back to sl
    sl = append(sl, 7)
    fmt.Println(sl)
    // finding the length of an array is done using the len() built in function
    l := len(sl2)
    fmt.Printf("%s: %d\n", "Len sl2", l)

    /*
        Now on to strings! They can be thought of like character slices, but they
        don't adhere to all the same rules or functions as regular slices. For example,
        strings CAN have their length evaluated using len(), however, they cannot be 
        concatenated by using append(), instead the plus operator (+) is used.
    */
    s := "Hello"
    slen := len(s)
    fmt.Printf("%s: %d\n", "string len: ", slen)
    // set s equal to itself concatenated with " World!"
    s += " World!"
    fmt.Println(s)
    // you can also lowercase or uppercase the entire string
    s = strings.ToUpper(s)
    fmt.Println(s)
    s = strings.ToLower(s)
    fmt.Println(s)
    /* 
        search and replace using strings.Replace() - takes 4 arguments
            - 1: string to search and replace on
            - 2: substring to search for
            - 3: string to replace it with (if substring is found)
            - 4: amount of occurences to replace. If -1, replace all found
    */
    s = strings.Replace(s, " world", "", -1)
    fmt.Println(s)
    /* 
        Search for a substring within a base string using strings.Contains()
    */
    if strings.Contains(s, "hello") {
        fmt.Println("s contains 'hello'")
    }
    /*
        strings can also be split into multiple parts based on any other
        string as a separating point. First, lets restore s back into 
        "Hello, World!", then separate it along the space character to get
        a slice of 2 elements
    */
    s = "Hello, World!"
    strSlice := strings.Split(s, " ")
    fmt.Println(strSlice)
}
