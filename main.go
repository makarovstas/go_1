package main

import "fmt"
import "github.com/beevik/ntp"

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println("Error while request time")
	} else {
		fmt.Printf("%s", time)
	}
}
