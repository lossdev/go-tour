package main

/*
    Gomodules allow us to separate and compartmentalize code that can 
    support or extend other parts of a codebase. All library imports are
    go modules - separate libraries that have exposed functions to call, and 
    serve some need in our code that we haven't explicitely written. In the 
    same fashion, we can make our own modules and use them.

    This is the driver file for our two other modules, firstModule and 
    secondModule. Those two other modules contain supporting code and 
    functions that can be imported into this file and called.
*/

import (
    /*
        We are making a module called 'example', where we'll be defining two new 
        submodules firstModule and secondModule, that contain their own methods
        and code. Each of those modules can be imported and used separately
    */
    "example/firstModule"
    "example/secondModule"
)

/*
    Executing 'go mod init example' in this directory will generate the go.mod
    and go.sum files. Then, a binary containing all the modules together can be
    built using 'go build':

    go build .
*/

func main() {
    firstModule.SayHello()
    secondModule.SayHello()
}
