package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//func for check input string and space character
func trimTrailingCrLf(s string) string {
	trimmed := s
	if strings.HasSuffix(trimmed, "\n") {
		trimmed = strings.TrimSuffix(trimmed, "\n")
	}
	if strings.HasSuffix(trimmed, "\r") {
		trimmed = strings.TrimSuffix(trimmed, "\r")
	}

	return trimmed
}

func main() {
	type Orang map[string]string //make map
	orang := make(Orang)

	stdinReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please enter the name:")
		if nama, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			orang["name"] = trimTrailingCrLf(nama)
			break
		}
	}

	for {
		fmt.Print("Please enter the address:")
		if alamat, err := stdinReader.ReadString('\n'); err != nil {
			fmt.Println("\nError: ", err)
		} else {
			orang["address"] = trimTrailingCrLf(alamat)
			break
		}
	}
	if data, err := json.Marshal(orang); err != nil {
		fmt.Println("\nError in serializing to JSON: ", err)
	} else {
		fmt.Println("\nJSON:", string(data))
	}
}
