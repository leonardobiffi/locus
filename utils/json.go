package utils

import (
	"fmt"

	json "github.com/json-iterator/go"
)

func PrintJson(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "  ")
	fmt.Println(string(s))
}
