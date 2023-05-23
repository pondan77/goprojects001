package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	//
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Invalid arguments!\nlottery filename count")
		return
	}

	filename := os.Args[1]
	count, err := strnconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot convert count to integer! count:", count)
		return
	}

	candidates, err := readCandidates(filename)
	winners := make([]string, count)
	for i := 0; i < count; i++ {
		n := rand.Intn(len(candidates))
		winners[i] = candidates[n]
		candidates = append(candidates[:n], candidates[n+1:]...)
	}

	fmt.Println("Winners are !!!")
	for _, winner := range winners {
		fmt.Println(winner)
	}
}
