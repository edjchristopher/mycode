/* RZFeeser | Alta3 Research
   Create a struct named person */

package main

import "fmt"

type Astro struct {
    name string
    age  int
    mission string
    isneeded bool
}

// this is our new struct
type nasaMission struct {
        people []Astro
        number int
        message string
}

func main() {

    ast1 := Astro{"Bob", 20, "Moon", true}

    ast2 := Astro{"Alice", 30, "Moon", true}

    ast3 := Astro{"Fred", 32, "Moon", false} 

    fmt.Println(ast1)

    fmt.Println(ast2)

    fmt.Println(ast3)

    // slice named people made up of Astro
    p := []Astro{ast1, ast2, ast3}

    // display the slice
    fmt.Println(p)

    // select data from the struct
    fmt.Println(p[2].mission)

    // initialize a nasaMission struct, "s"
    s := nasaMission{p, 3, "success"}

    // display "s" without fields
    fmt.Println(s)

    // display "s" with fields
    fmt.Printf("%+v", s)

}
