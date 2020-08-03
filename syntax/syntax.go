package main
/* this is a comment */
// this is also a comment
/*
    Every go program needs a 'package' declaration at the top. A 'package main' 
    is used when the program you are making is expected to be directly executed. 
    A program with any other package name is then implicitely a module (or library),
    supporting code that is not meant to be directly executed.
*/

/*
    An import() call specifies which libraries to use during runtime. Unlike
    most other languages, you CANNOT import a library and not use it (with the
    exception of using the _ operator preceding the import). Although many users
    don't like this paradigm, the Go language creators made this a requirement to
    keep binary sizes as low as possible, and prevent unnecessary code from being
    bundled with the necessary code.

    There are many, many modules that you can import into a Go program, both from
    the standard library as well as other users. Modules from other users (i.e those
    not directly made or maintained by the Go developers) will have some sort of 
    prefix attached to them that identifies where on the internet they're from. For
    example, many Go modules are hosted on Github, such as a simple stack library that
    I wrote and host on Github. That import would look like this:

    "github.com/lossdev/stack"

    However, the standard library imports don't have this URL-esque prefix. Some
    examples of standard library imports:

    "fmt"
    "net/http"
    "encoding/json"

    Simply knowing this distinction can help you quickly understand where yours
    or someone else's code is coming from, and where to find documentation for it.
*/
import (
    /*
        fmt is one of the standard input/output libraries in Go. fmt's methods can
        print to the console, take in input from the console, concatenate strings,
        and more. Libraries are expected to be imported one per line, enclosed in
        double quotations in an import() block.
    */
    "fmt"
)

/*
    Any 'package main' program will need to declare func main() somewhere. This
    is the entrypoint of our program.
*/
func main() {
    // Call fmt's Println() method to print the string argument we pass to it
    fmt.Println("Hello, World!")
    /*
        Imported libraries will expose certain functions/methods for the user to
        use. This is specified by calling [module].[method]
    */
}
