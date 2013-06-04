package main

import "fmt"
import "math"

func vaxMthRand (seed uint) uint {
 	return ((69069 * seed + 1) % uint(math.Pow(2,32))) 
}

func add(x int, y int) int {
    return x + y
}

func main() {
    var nextSeed, result uint
    nextSeed = 6
    for i := 0; i < 10; i++ {
    	nextSeed = vaxMthRand(nextSeed)
    	result = nextSeed % 36
    	fmt.Println(result)
    }
}