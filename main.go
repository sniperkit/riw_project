package main

import (
	"fmt"
	"time"
	"./parsers/cacm"
	"./requests"
)

func main() {
	start := time.Now()
	// Get a reversed dictionnary of relevant terms
	collection := cacm.NewCollection("./consignes/Data/CACM/")
	done := time.Now()
	elapsed := done.Sub(start)
	fmt.Printf("Index computed in %f seconds\n", elapsed.Seconds())

	start = time.Now()
	request := requests.NewBinaryRequest("topologically & 36-bit | chicken", collection.Index)
	userOutput := requests.NewUserOutput(request.Output, request.DocsScore)
	userOutput.Print()
	done = time.Now()
	elapsed = done.Sub(start)
	fmt.Printf("Result computed in %f seconds\n", elapsed.Seconds())
}