package app_configs

import (
	"fmt"
	"os"
)

var (
	Port                     = fmt.Sprintf(":%s", os.Getenv("PORT"))
	CtxAcceptList [][]string = [][]string{
		{"html"},
		{"text/html"},
		{"json", "text"},
		{"application/json"},
		{"text/plain", "application/json"},
		{"image/png"},
		{"png"},
	}
)
