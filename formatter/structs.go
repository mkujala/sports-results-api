package formatter

import (
	"encoding/json"
	"fmt"
	"log"
)

// PrettyPrint is used for debugging purpuses to get pretty prints of structs to console
func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
		return
	}
	log.Fatal(err)
}
