package main

import "fmt"

type flt func(int) bool
type slice_split func([] int)([] int, [] int)

func isOdd(integer int) bool {  // check if integer is odd
    if integer % 2 == 0 {
        return false
    }

    return true
}

func isBiggerThan4(integer int) bool {  // check if integer is greater than 4 
    if integer > 4 {
        return true
    }
    return false
}

func filter_factory(f flt) slice_split {    // split the slice on basis of func
    return func(s[] int)(yes, no[] int) {
        for _, val := range s {
            if f(val) {
                yes = append(yes, val)
            } else {
                no = append(no, val)
            }
        }
        return
    }
}

func main() {
    s := [] int {1, 2, 3, 4, 5, 7}
    fmt.Println("s = ", s)
    odd_even_function := filter_factory(isOdd)
    odd,even := odd_even_function(s)
    fmt.Println("odd = ", odd)
    fmt.Println("even = ", even)
    //separate those that are bigger than 4 and those that are not.
    bigger,smaller := filter_factory(isBiggerThan4)(s)
    fmt.Println("Bigger than 4: ", bigger)
    fmt.Println("Smaller than or equal to 4: ", smaller)
}