package main

import (
	"github.com/PARKNAMSU/user-management/src"
	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load()
)

func main() {
	src.AppInit()
}
