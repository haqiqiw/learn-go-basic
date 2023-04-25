package main

import (
	"flag"
	"fmt"
)

// go run flag.go -host=127.0.0.1 -username=wawan -password=knalpot
func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
