package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("token not provided")
		return
	} else if len(args) > 2 {
		fmt.Println("too many arguments")
		return
	}

	token := args[1]
	tokenParts := strings.Split(token, ".")

	if len(tokenParts) != 3 {
		fmt.Println("invalid token")
		return
	}

	header, err := base64.RawURLEncoding.DecodeString(tokenParts[0])
	if err != nil {
		fmt.Println("invalid token")
		return
	}

	payload, err := base64.RawURLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		fmt.Println("invalid token")
		return
	}

	var prettyHeader bytes.Buffer
	err = json.Indent(&prettyHeader, header, "", "  ")
	if err != nil {
		fmt.Println("invalid token")
		return
	}

	var prettyPayload bytes.Buffer
	err = json.Indent(&prettyPayload, payload, "", "  ")
	if err != nil {
		fmt.Println("invalid token")
		return
	}

	fmt.Println()
	fmt.Println("Header:")
	fmt.Println(prettyHeader.String())

	fmt.Println()
	fmt.Println("Paylod:")
	fmt.Println(prettyPayload.String())
}
