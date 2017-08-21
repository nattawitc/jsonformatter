package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	for {
		data := make(map[string]interface{})
		if err := dec.Decode(&data); err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintln(os.Stderr, err)
				return
			}
		}
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println(string(b))
	}
}
