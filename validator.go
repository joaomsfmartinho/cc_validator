package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Function that calculates the checksum using the Luhn algorithm
func checksum(number int64) int64 {
	var luhn int64

	for i := 0; number > 0; i++ {
		cur := number % 10
		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}

// Function that validates the credit card number using the Luhn algorithm
func luhnAlgorithm(ccNum string) bool {
	// Parse the credit card number string into an integer
	number, err := strconv.ParseInt(ccNum, 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return (number%10+checksum(number/10))%10 == 0
}

// Handle the HTML form GET
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving request: %s\n", r.URL.Path)
	io.WriteString(w, "<html><head><title>CC Validator</title></head><body><h1>CC Validator</h1>"+
		"<p>Enter a credit card number to validate</p>"+
		"<form action=\"/hello/\" method=\"post\"><input type=\"text\" name=\"ccnum\">"+
		"<input type=\"submit\" value=\"Submit\"></form></body></html>")
}

// Handle the HTML form POST
func handlePost(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving request: %s\n", r.URL.Path)
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	ccNum := r.FormValue("ccnum")
	if luhnAlgorithm(ccNum) {
		io.WriteString(w, "<html><head><title>CC Validator</title></head><body><h1>CC Validator</h1>"+
			"<p>"+ccNum+" is a valid credit card number</p></body></html>")
	} else {
		io.WriteString(w, "<html><head><title>CC Validator</title></head><body><h1>CC Validator</h1>"+
			"<p>"+ccNum+" is not a valid credit card number</p></body></html>")
	}
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello/", handlePost)
	http.ListenAndServe(":3333", nil)
}
