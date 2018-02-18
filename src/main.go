package main

import (
	//"compress/zlib"
	"encoding/base64"
	"fmt"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// https://pastebin.com/raw/ZBQHAF2U

func rawpaste(url string) string {
	return strings.Replace(url, "pastebin.com/", "pastebin.com/raw/", 1)
}

func replace(s string) string {
	s = strings.Replace(s, "-", "+", -1)
	return strings.Replace(s, "_", "/", -1)
}

func DecodeB64(encoded string) string {
	data, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		log.Fatal("error:", err)
	}
	decoded := string(data)
	print(decoded)
	return decoded
}

func main() {
	fmt.Printf("Hello World\n")

	var pastebin string

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	} else {
		if strings.Contains(os.Args[1], "pastebin.com/raw/") == true {
			fmt.Print("Already a raw paste URL, skipping\n")
			pastebin = os.Args[1]
		} else {
			if strings.Contains(os.Args[1], "pastebin.com") == true {
				fmt.Print("This is a Pastebin URL, adding 'raw' \n")
				pastebin = rawpaste(os.Args[1])
				//fmt.Print(pastebin)
			}
		}
	}

	response, err := http.Get(pastebin)

	fmt.Println("response:", response, "\n")
	fmt.Println("response type:", reflect.TypeOf(response), "\n")
	fmt.Println("response body:", response.Body, "\n")

	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		responseData, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal("error:", err)
		}

		responseString := string(responseData)

		fmt.Println(replace(responseString))
		fmt.Println(DecodeB64(responseString))
	}
}
