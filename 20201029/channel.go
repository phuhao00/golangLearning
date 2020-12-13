
// Go program to illustrate how to
// find the capacity of the channel

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//Length of the Channel: In channel, you can find the length of the channel using len() function. Here,
//the length indicates the number of value queued in the channel buffer.
func main() {

	fmt.Println(rand.Intn(1))
	return

	// Creating a channel
	// Using make() function
	mychnl := make(chan string, 5)
	mychnl <- "GFG"
	mychnl <- "gfg"
	mychnl <- "Geeks"
	mychnl <- "GeeksforGeeks"

	// Finding the capacity of the channel
	// Using cap() function
	fmt.Println("Capacity of the channel is: ", cap(mychnl),len(mychnl))
	sync.Pool{}
}

