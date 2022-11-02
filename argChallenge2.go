/* Alta3 Research | RZFeeser
   Reading positional arguments from the command line*/
   
package main

import (
    "fmt"
    "os"
)

func main() {

    // os.Args provides access to raw CLI arguments
    // argsWithProg := os.Args         // includes program name
    
    // argsWithoutProg := os.Args[1:]  // does not include program name

    // arg := os.Args[3]               // select the 3rd argument
                                    // remember: the 1st argument is the name
                                    //           of the program

    // last_arg := os.Args[len(os.Args)-1]         // select the last argument         // select the last argument

    // display the results
    for i, a := range os.Args[1:] {                 // loop across args
        fmt.Printf("Arg %d is %s\n", i+1, a)        // use a format string
    }

}

