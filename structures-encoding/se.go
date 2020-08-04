package main

import (
    /*
        Go is particularly good at serializing JSON data. Golang structures
        also coincide very naturally with the structure of a Golang struct. 
        Golang is a widely used language for HTTP servers and API's, which 
        commonly use JSON as a data description language - hence, why 
        the encoding/json library is a valuable tool to learn.
    */
    "encoding/json"
    "fmt"
    "log"
)


// structs hold multiple types of data, and can be turned into JSON data

type Foo struct {
    Bar string
    Baz string
}

func main() {
    /* 
        instantiate a copy of our struct - two ways to do it: you can not specify 
        which variables you're filling as long as you fill them in the order they're
        declared in the struct, or specify which variable you're filling as you 
        fill it. Both are below:
    */
    f := Foo{"string1", "string2"}
    f = Foo{Bar: "string1", Baz: "string2"}
    /*
        Now, we can serialize our structure into JSON data using the encoding/json
        library's Marshal() method
    */
    j, err := json.Marshal(f)
    if err != nil {
        log.Println(err)
        return
    }
    /*
        Serialized json has a []byte (or byte slice) data type - therefore, you
        can easily typecast it to a string and print it
    */
    fmt.Println(string(j))
    // Now, let's try the opposite and turn our json data back into a struct
    // This is accomplished by calling json.Unmarshal()
    var e Foo
    err = json.Unmarshal(j, &e)
    fmt.Println(e)
}
