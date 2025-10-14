package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetPrefix("jwt-dump: ")
	log.SetFlags(log.LstdFlags &^ (log.Ldate | log.Ltime))

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter token:")
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatalln("could not read input -", err)
	}

	token := scanner.Text()
	tokenParts := strings.Split(token, ".")

	if len(tokenParts) != 3 {
		log.Fatalln("invalid token")
	}

	header, err := base64.RawURLEncoding.DecodeString(tokenParts[0])
	if err != nil {
		log.Fatalln("could not decode header")
	}

	payload, err := base64.RawURLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		log.Fatalln("could not decode payload")
	}

	var prettyHeader bytes.Buffer
	err = json.Indent(&prettyHeader, header, "", "  ")
	if err != nil {
		log.Fatalln("invalid header json")
	}

	var prettyPayload bytes.Buffer
	err = json.Indent(&prettyPayload, payload, "", "  ")
	if err != nil {
		log.Fatalln("invalid payload json")
	}

	fmt.Println()
	fmt.Println("Header:")
	fmt.Println(prettyHeader.String())

	fmt.Println()
	fmt.Println("Payload:")
	fmt.Println(prettyPayload.String())
}
