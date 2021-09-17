package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("MONGO_INITDB_ROOT_USERNAME"))
}
