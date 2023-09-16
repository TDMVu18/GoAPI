package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load(".env")
	uri := os.Getenv("MG_CONNECT")
	fmt.Println(uri)
}
